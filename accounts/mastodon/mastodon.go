// Package mastodon is a Mastodon account for Telephant.
package mastodon

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/mattn/go-mastodon"

	"github.com/muesli/telephant/accounts"
)

const (
	initialFeedCount          = 40
	initialNotificationsCount = 40
)

// Account is a Mastodon account for Telephant.
type Account struct {
	client *mastodon.Client
	config *mastodon.Config
	self   *mastodon.Account

	evchan  chan interface{}
	SigChan chan bool
}

// NewAccount returns a new Mastodon account.
func NewAccount(instance, token, clientID, clientSecret string) *Account {
	mconfig := &mastodon.Config{
		Server:       instance,
		AccessToken:  token,
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	return &Account{
		config: mconfig,
		client: mastodon.NewClient(mconfig),
	}
}

// RegisterAccount registers the app with an instance and retrieves an
// authentication URI.
func RegisterAccount(instance string) (*Account, string, string, error) {
	app, err := mastodon.RegisterApp(context.Background(), &mastodon.AppConfig{
		Server:     instance,
		ClientName: "Telephant",
		Scopes:     "read write follow post",
		Website:    "https://github.com/muesli/telephant",
	})
	if err != nil {
		return nil, "", "", err
	}

	a := NewAccount(instance, "", app.ClientID, app.ClientSecret)

	return a, app.AuthURI, app.RedirectURI, nil
}

// Authenticate finishes the authentication and retrieves an accesstoken from
// the instance, which we'll use for future logins.
func (mod *Account) Authenticate(code, redirectURI string) (string, string, string, string, error) {
	if redirectURI == "" {
		redirectURI = "urn:ietf:wg:oauth:2.0:oob"
	}
	err := mod.client.AuthenticateToken(context.Background(), code, redirectURI)
	if err != nil {
		return "", "", "", "", err
	}

	return mod.config.Server, mod.config.AccessToken, mod.config.ClientID, mod.config.ClientSecret, nil
}

// Run executes the account's event loop.
func (mod *Account) Run(eventChan chan interface{}) error {
	if mod.config.AccessToken == "" {
		return errors.New("no accesstoken provided")
	}

	mod.evchan = eventChan

	var err error
	mod.self, err = mod.client.GetAccountCurrentUser(context.Background())
	if err != nil {
		ev := accounts.ErrorEvent{
			Message:  err.Error(),
			Internal: false,
		}
		mod.evchan <- ev
		return err
	}

	ev := accounts.LoginEvent{
		Username:      mod.self.Acct,
		Name:          mod.self.DisplayName,
		Avatar:        mod.self.Avatar,
		ProfileURL:    mod.self.URL,
		ProfileID:     string(mod.self.ID),
		Posts:         mod.self.StatusesCount,
		FollowCount:   mod.self.FollowingCount,
		FollowerCount: mod.self.FollowersCount,
		PostSizeLimit: 500, // FIXME: retrieve from API, once possible
	}
	if strings.TrimSpace(ev.Name) == "" {
		ev.Name = mod.self.Username
	}
	mod.evchan <- ev

	// seed feeds initially
	nn, err := mod.client.GetNotifications(context.Background(), &mastodon.Pagination{
		Limit: initialNotificationsCount,
	})
	if err != nil {
		ev := accounts.ErrorEvent{
			Message:  err.Error(),
			Internal: false,
		}
		mod.evchan <- ev
		return err
	}
	for i := len(nn) - 1; i >= 0; i-- {
		mod.handleNotification(nn[i], false)
	}

	tt, err := mod.client.GetTimelineHome(context.Background(), &mastodon.Pagination{
		Limit: initialFeedCount,
	})
	if err != nil {
		ev := accounts.ErrorEvent{
			Message:  err.Error(),
			Internal: false,
		}
		mod.evchan <- ev
		return err
	}
	for i := len(tt) - 1; i >= 0; i-- {
		mod.evchan <- mod.handleStatus(tt[i])
	}

	go mod.handleStream()
	return nil
}

// Panes returns the panes this Mastodon account offers
func (mod *Account) Panes() []accounts.Pane {
	ll, err := mod.client.GetLists(context.Background())
	if err != nil {
		fmt.Println("Error retrieving lists:", err)
	}

	p := []accounts.Pane{}
	for _, list := range ll {
		// take a copy of the ID here, so the closure accesses an absolute value
		id := string(list.ID)
		p = append(p, accounts.Pane{
			Title: "List: " + list.Title,
			Stream: func(ch chan interface{}) error {
				return mod.List(id, ch)
			},
		})
	}

	return p
}

// Logo returns the Mastodon logo.
func (mod *Account) Logo() string {
	return "mastodon.svg"
}

// Post posts a new status
func (mod *Account) Post(message string, attachments []string) error {
	t := &mastodon.Toot{
		Status: message,
	}
	for _, v := range attachments {
		t.MediaIDs = append(t.MediaIDs, mastodon.ID(v))
	}

	_, err := mod.client.PostStatus(context.Background(), t)
	return err
}

// Reply posts a new reply-status
func (mod *Account) Reply(replyid string, message string, attachments []string) error {
	t := &mastodon.Toot{
		Status:      message,
		InReplyToID: mastodon.ID(replyid),
	}
	for _, v := range attachments {
		t.MediaIDs = append(t.MediaIDs, mastodon.ID(v))
	}

	_, err := mod.client.PostStatus(context.Background(), t)
	return err
}

// UploadAttachment uploads an attachment to Mastodon
func (mod *Account) UploadAttachment(url string) {
	go func() {
		a, err := mod.client.UploadMedia(context.Background(), url)
		if err != nil {
			ev := accounts.ErrorEvent{
				Message:  err.Error(),
				Internal: false,
			}
			mod.evchan <- ev
			return
		}

		mod.evchan <- accounts.Media{
			ID:      string(a.ID),
			Preview: a.PreviewURL,
		}
	}()
}

// DeletePost deletes a post
func (mod *Account) DeletePost(id string) error {
	err := mod.client.DeleteStatus(context.Background(), mastodon.ID(id))
	return err
}

// Share boosts a post
func (mod *Account) Share(id string) error {
	_, err := mod.client.Reblog(context.Background(), mastodon.ID(id))
	return err
}

// Unshare deletes a boost for a post
func (mod *Account) Unshare(id string) error {
	_, err := mod.client.Unreblog(context.Background(), mastodon.ID(id))
	return err
}

// Like favourites a post
func (mod *Account) Like(id string) error {
	_, err := mod.client.Favourite(context.Background(), mastodon.ID(id))
	return err
}

// Unlike un-favourites a post
func (mod *Account) Unlike(id string) error {
	_, err := mod.client.Unfavourite(context.Background(), mastodon.ID(id))
	return err
}

// Follow follows another user
func (mod *Account) Follow(id string) error {
	_, err := mod.client.AccountFollow(context.Background(), mastodon.ID(id))
	return err
}

// Unfollow unfollows another user
func (mod *Account) Unfollow(id string) error {
	_, err := mod.client.AccountUnfollow(context.Background(), mastodon.ID(id))
	return err
}

// Tag starts a hashtag stream
func (mod *Account) Tag(token string, ch chan interface{}) error {
	tt, err := mod.client.GetTimelineHashtag(context.Background(), token, false, &mastodon.Pagination{
		Limit: initialFeedCount,
	})
	if err != nil {
		ev := accounts.ErrorEvent{
			Message:  err.Error(),
			Internal: false,
		}
		mod.evchan <- ev
		return err
	}
	for i := len(tt) - 1; i >= 0; i-- {
		ch <- mod.handleStatus(tt[i])
	}

	s, err := mod.client.StreamingHashtag(context.Background(), token, false)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-mod.SigChan:
				return
			case item := <-s:
				mod.handleStreamEvent(item, ch)
			}
		}
	}()

	return nil
}

// List starts a list stream
func (mod *Account) List(id string, ch chan interface{}) error {
	tt, err := mod.client.GetTimelineList(context.Background(), mastodon.ID(id), &mastodon.Pagination{
		Limit: initialFeedCount,
	})
	if err != nil {
		ev := accounts.ErrorEvent{
			Message:  err.Error(),
			Internal: false,
		}
		mod.evchan <- ev
		return err
	}
	for i := len(tt) - 1; i >= 0; i-- {
		ch <- mod.handleStatus(tt[i])
	}

	s, err := mod.client.StreamingList(context.Background(), mastodon.ID(id))
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-mod.SigChan:
				return
			case item := <-s:
				mod.handleStreamEvent(item, ch)
			}
		}
	}()

	return nil
}

// LoadConversation loads a message conversation
func (mod *Account) LoadConversation(id string) ([]accounts.MessageEvent, error) {
	var r []accounts.MessageEvent

	status, err := mod.client.GetStatus(context.Background(), mastodon.ID(id))
	if err != nil {
		return r, err
	}
	contexts, err := mod.client.GetStatusContext(context.Background(), mastodon.ID(id))
	if err != nil {
		return r, err
	}

	fmt.Printf("Found %d ancestors and %d descendants\n", len(contexts.Ancestors), len(contexts.Descendants))
	for _, m := range contexts.Ancestors {
		r = append(r, mod.handleStatus(m))
	}

	r = append(r, mod.handleStatus(status))

	for _, m := range contexts.Descendants {
		r = append(r, mod.handleStatus(m))
	}

	return r, nil
}

// LoadAccount loads a profile's information.
func (mod *Account) LoadAccount(id string) (accounts.ProfileEvent, []accounts.MessageEvent, error) {
	var p accounts.ProfileEvent
	var r []accounts.MessageEvent

	a, err := mod.client.GetAccount(context.Background(), mastodon.ID(id))
	if err != nil {
		return p, r, err
	}

	p = accounts.ProfileEvent{
		Username:      a.Acct,
		Name:          a.DisplayName,
		Avatar:        a.Avatar,
		ProfileURL:    a.URL,
		ProfileID:     string(a.ID),
		Posts:         a.StatusesCount,
		FollowCount:   a.FollowingCount,
		FollowerCount: a.FollowersCount,
	}
	if strings.TrimSpace(p.Name) == "" {
		p.Name = a.Username
	}

	f, err := mod.client.GetAccountRelationships(context.Background(), []string{id})
	if err != nil {
		return p, r, err
	}
	if len(f) > 0 {
		p.Following = f[0].Following
		p.FollowedBy = f[0].FollowedBy
	}

	tt, err := mod.client.GetAccountStatuses(context.Background(), mastodon.ID(id), &mastodon.Pagination{
		Limit: 40,
	})
	if err != nil {
		return p, r, err
	}

	for _, m := range tt {
		r = append(r, mod.handleStatus(m))
	}

	return p, r, nil
}

func (mod *Account) PostLimitCount(body string) int {
	r := regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,63}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)
	body = r.ReplaceAllString(body, strings.Repeat(" ", 23))

	return len(body)
}

// parseBody cleans up a post's content.
func parseBody(s *mastodon.Status) string {
	body := s.Content

	// hide invisible message parts
	r := regexp.MustCompile("<span class=\"invisible\">(.[^<]*)</span>")
	body = r.ReplaceAllString(body, "")

	// expand ellipsis
	r = regexp.MustCompile("<span class=\"ellipsis\">(.[^<]*)</span>")
	body = r.ReplaceAllString(body, "$1...")

	for _, u := range s.Mentions {
		body = strings.Replace(body, u.URL, fmt.Sprintf("telephant://mastodon/user/%s", u.ID), -1)
	}
	for _, t := range s.Tags {
		r = regexp.MustCompile("<a href=\"(.[^\"]*)/tags/" + t.Name + "\"")
		body = r.ReplaceAllString(body, "<a href=\"telephant://mastodon/tag/"+t.Name+"\"")
	}
	return body

	/*
		for _, u := range ents.Urls {
			r := fmt.Sprintf("<a style=\"text-decoration: none; color: orange;\" href=\"%s\">%s</a>", u.Expanded_url, u.Display_url)
			ev.Post.Body = strings.Replace(ev.Post.Body, u.Url, r, -1)
		}
		for _, media := range ents.Media {
			ev.Media = append(ev.Media, media.Media_url_https)
			ev.Post.Body = strings.Replace(ev.Post.Body, media.Url, "", -1)
			// FIXME:
			break
		}
	*/
}

// handleNotification handles incoming notification events.
func (mod *Account) handleNotification(n *mastodon.Notification, notify bool) {
	var ev accounts.MessageEvent
	if n.Status != nil {
		ev = accounts.MessageEvent{
			Account:      "mastodon",
			Name:         "post",
			Notification: true,
			Notify:       notify,

			Post: accounts.Post{
				MessageID:  string(n.Status.ID),
				Body:       parseBody(n.Status),
				Sensitive:  n.Status.Sensitive,
				Warning:    n.Status.SpoilerText,
				Author:     n.Account.Acct,
				AuthorName: n.Account.DisplayName,
				AuthorURL:  n.Account.URL,
				AuthorID:   string(n.Account.ID),
				Avatar:     n.Account.Avatar,
				CreatedAt:  n.CreatedAt,
				URL:        n.Status.URL,
			},
		}
		ev.Post.Liked, _ = n.Status.Favourited.(bool)
		ev.Post.Shared, _ = n.Status.Reblogged.(bool)
		if strings.TrimSpace(ev.Post.AuthorName) == "" {
			ev.Post.AuthorName = n.Account.Username
		}

		for _, v := range n.Status.MediaAttachments {
			ev.Media = append(ev.Media, accounts.Media{
				Preview: v.PreviewURL,
				URL:     v.URL,
			})
		}
	}

	switch n.Type {
	case "mention":
		if n.Status.InReplyToID != nil {
			ev.Mention = true
			ev.Post.ReplyToAuthor = n.Status.InReplyToAccountID.(string)
			ev.Post.ReplyToID = n.Status.InReplyToID.(string)
		}

	case "reblog":
		ev.Forward = true
		ev.Post.MessageID = n.Account.Acct + "-share-" + ev.Post.MessageID
		ev.Post.Author = n.Status.Account.Acct
		ev.Post.AuthorName = n.Status.Account.DisplayName
		ev.Post.AuthorURL = n.Status.Account.URL
		ev.Post.AuthorID = string(n.Status.Account.ID)
		ev.Post.Avatar = n.Status.Account.Avatar
		ev.Post.Actor = n.Account.Acct
		ev.Post.ActorName = n.Account.DisplayName
		ev.Post.ActorID = string(n.Account.ID)

		if strings.TrimSpace(ev.Post.AuthorName) == "" {
			ev.Post.AuthorName = n.Status.Account.Username
		}
		if strings.TrimSpace(ev.Post.ActorName) == "" {
			ev.Post.ActorName = n.Account.Username
		}

		ev.Post.Body = parseBody(n.Status)

	case "favourite":
		ev.Like = true
		ev.Post.MessageID = n.Account.Acct + "-like-" + ev.Post.MessageID
		ev.Post.Author = n.Status.Account.Acct
		ev.Post.AuthorName = n.Status.Account.DisplayName
		ev.Post.AuthorURL = n.Status.Account.URL
		ev.Post.AuthorID = string(n.Status.Account.ID)
		ev.Post.Avatar = n.Status.Account.Avatar
		ev.Post.Actor = n.Account.Acct
		ev.Post.ActorName = n.Account.DisplayName
		ev.Post.ActorID = string(n.Account.ID)

		if strings.TrimSpace(ev.Post.AuthorName) == "" {
			ev.Post.AuthorName = n.Status.Account.Username
		}
		if strings.TrimSpace(ev.Post.ActorName) == "" {
			ev.Post.ActorName = n.Account.Username
		}

		ev.Post.Body = parseBody(n.Status)

	case "follow":
		ev = accounts.MessageEvent{
			Account:      "mastodon",
			Name:         "follow",
			Notification: true,
			Notify:       notify,
			Followed:     true,
			Follow: accounts.Follow{
				Account:    n.Account.Acct,
				Name:       n.Account.DisplayName,
				Avatar:     n.Account.Avatar,
				ProfileURL: n.Account.URL,
				ProfileID:  string(n.Account.ID),
			},
		}

		f, _ := mod.client.GetAccountRelationships(context.Background(), []string{string(n.Account.ID)})
		if len(f) > 0 {
			ev.Follow.Following = f[0].Following
			ev.Follow.FollowedBy = f[0].FollowedBy
		}

		if strings.TrimSpace(ev.Follow.Name) == "" {
			ev.Follow.Name = n.Account.Username
		}

	default:
		fmt.Println("Unknown type:", n.Type)
		return
	}

	mod.evchan <- ev
}

// handleStatus handles incoming status updates.
func (mod *Account) handleStatus(s *mastodon.Status) accounts.MessageEvent {
	ev := accounts.MessageEvent{
		Account: "mastodon",
		Name:    "post",
		Post: accounts.Post{
			MessageID:  string(s.ID),
			Body:       parseBody(s),
			Sensitive:  s.Sensitive,
			Warning:    s.SpoilerText,
			Author:     s.Account.Acct,
			AuthorName: s.Account.DisplayName,
			AuthorURL:  s.Account.URL,
			AuthorID:   string(s.Account.ID),
			Avatar:     s.Account.Avatar,
			CreatedAt:  s.CreatedAt,
			URL:        s.URL,
		},
	}
	ev.Post.Liked, _ = s.Favourited.(bool)
	ev.Post.Shared, _ = s.Reblogged.(bool)
	if strings.TrimSpace(ev.Post.AuthorName) == "" {
		ev.Post.AuthorName = s.Account.Username
	}

	for _, v := range s.MediaAttachments {
		ev.Media = append(ev.Media, accounts.Media{
			Preview: v.PreviewURL,
			URL:     v.URL,
		})
	}

	if s.Reblog != nil {
		ev.Forward = true

		for _, v := range s.Reblog.MediaAttachments {
			ev.Media = append(ev.Media, accounts.Media{
				Preview: v.PreviewURL,
				URL:     v.URL,
			})
		}

		ev.Post.URL = s.Reblog.URL
		ev.Post.Author = s.Reblog.Account.Acct
		ev.Post.AuthorName = s.Reblog.Account.DisplayName
		ev.Post.AuthorURL = s.Reblog.Account.URL
		ev.Post.AuthorID = string(s.Reblog.Account.ID)
		ev.Post.Avatar = s.Reblog.Account.Avatar
		ev.Post.Actor = s.Account.Acct
		ev.Post.ActorName = s.Account.DisplayName
		ev.Post.ActorID = string(s.Account.ID)

		ev.Post.Liked, _ = s.Reblog.Favourited.(bool)
		ev.Post.Shared, _ = s.Reblog.Reblogged.(bool)

		if strings.TrimSpace(ev.Post.AuthorName) == "" {
			ev.Post.AuthorName = s.Reblog.Account.Username
		}
		if strings.TrimSpace(ev.Post.ActorName) == "" {
			ev.Post.ActorName = s.Account.Username
		}

		ev.Post.Body = parseBody(s.Reblog)
	}

	return ev
}

// handleStreamEvent handles incoming events and dispatches them to the correct
// handler.
func (mod *Account) handleStreamEvent(item interface{}, ch chan interface{}) {
	// spw := &spew.ConfigState{Indent: "  ", DisableCapacities: true, DisablePointerAddresses: true}
	// log.Println("Message received:", spw.Sdump(item))

	switch e := item.(type) {
	case *mastodon.NotificationEvent:
		mod.handleNotification(e.Notification, true)

	case *mastodon.UpdateEvent:
		ch <- mod.handleStatus(e.Status)

	case *mastodon.DeleteEvent:
		mod.evchan <- accounts.DeleteEvent{
			ID: string(e.ID),
		}

	default:
		fmt.Printf("Unknown event: %+v\n", item)
	}
}

// handleStream handles all connected Mastodon API streams.
func (mod *Account) handleStream() {
	timeline, err := mod.client.StreamingUser(context.Background())
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-mod.SigChan:
			return
		case item := <-timeline:
			mod.handleStreamEvent(item, mod.evchan)
		}
	}
}
