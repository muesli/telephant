// Package mastodon is a Mastodon account for Chirp.
package mastodon

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/davecgh/go-spew/spew"
	"github.com/mattn/go-mastodon"

	"github.com/muesli/chirp/accounts"
)

const (
	initialFeedCount          = 200
	initialNotificationsCount = 50
)

// Account is a Mastodon account for Chirp.
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

func RegisterAccount(instance string) (*Account, string, string, error) {
	app, err := mastodon.RegisterApp(context.Background(), &mastodon.AppConfig{
		Server:     instance,
		ClientName: "Telephant",
		Scopes:     "read write follow post",
		Website:    "",
	})
	if err != nil {
		return nil, "", "", err
	}

	a := NewAccount(instance, "", app.ClientID, app.ClientSecret)

	return a, app.AuthURI, app.RedirectURI, nil
}

func (mod *Account) Authenticate(code string) (string, string, string, string, error) {
	err := mod.client.AuthenticateToken(context.Background(), code, "urn:ietf:wg:oauth:2.0:oob")
	if err != nil {
		return "", "", "", "", err
	}

	return mod.config.Server, mod.config.AccessToken, mod.config.ClientID, mod.config.ClientSecret, nil
}

// Run executes the account's event loop.
func (mod *Account) Run(eventChan chan interface{}) {
	mod.evchan = eventChan

	if mod.config.AccessToken == "" {
		return
	}

	var err error
	mod.self, err = mod.client.GetAccountCurrentUser(context.Background())
	if err != nil {
		panic(err)
	}

	ev := accounts.LoginEvent{
		Username:   mod.self.Username,
		Name:       mod.self.DisplayName,
		Avatar:     mod.self.Avatar,
		ProfileURL: mod.self.URL,
	}
	mod.evchan <- ev

	// FIXME: retrieve initial feed
	mod.handleStream()
}

func (mod *Account) Logo() string {
	return "mastodon.svg"
}

// Post posts a new status
func (mod *Account) Post(message string) error {
	_, err := mod.client.PostStatus(context.Background(), &mastodon.Toot{
		Status: message,
	})
	return err
}

// Reply posts a new reply-status
func (mod *Account) Reply(replyid string, message string) error {
	_, err := mod.client.PostStatus(context.Background(), &mastodon.Toot{
		Status:      message,
		InReplyToID: mastodon.ID(replyid),
	})
	return err
}

// Share boosts a post
func (mod *Account) Share(id string) error {
	_, err := mod.client.Reblog(context.Background(), mastodon.ID(id))
	return err
}

// Like favourites a post
func (mod *Account) Like(id string) error {
	_, err := mod.client.Favourite(context.Background(), mastodon.ID(id))
	return err
}

func handleRetweetStatus(status string) string {
	/*
		if strings.HasPrefix(status, "RT ") && strings.Count(status, " ") >= 2 {
			return strings.Join(strings.Split(status, " ")[2:], " ")
		}
	*/

	return status
}

func handleReplyStatus(status string) string {
	/*
		if strings.HasPrefix(status, "@") && strings.Index(status, " ") > 0 {
			return status[strings.Index(status, " "):]
		}
	*/

	return status
}

func parseBody(body string) string {
	r := regexp.MustCompile("<span class=\"invisible\">(.[^<]*)</span>")
	body = r.ReplaceAllString(body, "")

	r = regexp.MustCompile("<span class=\"ellipsis\">(.[^<]*)</span>")
	return r.ReplaceAllString(body, "$1...")

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

func (mod *Account) handleStreamEvent(item interface{}) {
	spw := &spew.ConfigState{Indent: "  ", DisableCapacities: true, DisablePointerAddresses: true}
	log.Println("Message received:", spw.Sdump(item))

	switch status := item.(type) {
	case *mastodon.NotificationEvent:
		var ev accounts.MessageEvent
		if status.Notification.Status != nil {
			ev = accounts.MessageEvent{
				Account:      "mastodon",
				Name:         "post",
				Notification: true,

				Post: accounts.Post{
					MessageID:  string(status.Notification.Status.ID),
					Body:       status.Notification.Status.Content,
					Author:     status.Notification.Account.Username,
					AuthorName: status.Notification.Account.DisplayName,
					AuthorURL:  status.Notification.Account.URL,
					Avatar:     status.Notification.Account.Avatar,
					CreatedAt:  time.Now(),
					URL:        status.Notification.Status.URL,
				},
			}

			for _, v := range status.Notification.Status.MediaAttachments {
				ev.Media = append(ev.Media, v.PreviewURL)
			}
		}

		switch status.Notification.Type {
		case "mention":
			if status.Notification.Status.InReplyToID != nil {
				ev.Mention = true
				ev.Post.ReplyToAuthor = status.Notification.Status.InReplyToAccountID.(string)
				ev.Post.ReplyToID = status.Notification.Status.InReplyToID.(string)
			}

		case "reblog":
			ev.Forward = true
			ev.Post.Author = status.Notification.Status.Account.Username
			ev.Post.AuthorName = status.Notification.Status.Account.DisplayName
			ev.Post.AuthorURL = status.Notification.Status.Account.URL
			// ev.Post.Avatar = status.Notification.Status.Account.Avatar
			ev.Post.Actor = status.Notification.Account.Username
			ev.Post.ActorName = status.Notification.Account.DisplayName

		case "favourite":
			ev.Like = true

			ev.Post.Author = status.Notification.Status.Account.Username
			ev.Post.AuthorName = status.Notification.Status.Account.DisplayName
			ev.Post.AuthorURL = status.Notification.Status.Account.URL
			// ev.Post.Avatar = status.Notification.Status.Account.Avatar
			ev.Post.Actor = status.Notification.Account.Username
			ev.Post.ActorName = status.Notification.Account.DisplayName

		default:
			fmt.Println("Unknown type:", status.Notification.Type)
			return
		}

		mod.evchan <- ev

	case *mastodon.UpdateEvent:
		ev := accounts.MessageEvent{
			Account: "mastodon",
			Name:    "post",
			Post: accounts.Post{
				MessageID:  string(status.Status.ID),
				Body:       status.Status.Content,
				Author:     status.Status.Account.Acct,
				AuthorName: status.Status.Account.DisplayName,
				AuthorURL:  status.Status.Account.URL,
				Avatar:     status.Status.Account.Avatar,
				CreatedAt:  time.Now(),
				URL:        status.Status.URL,
			},
		}

		for _, v := range status.Status.MediaAttachments {
			ev.Media = append(ev.Media, v.PreviewURL)
		}

		if status.Status.Reblog != nil {
			ev.Forward = true

			for _, v := range status.Status.Reblog.MediaAttachments {
				ev.Media = append(ev.Media, v.PreviewURL)
			}

			ev.Post.URL = status.Status.Reblog.URL
			ev.Post.Author = status.Status.Reblog.Account.Username
			ev.Post.AuthorName = status.Status.Reblog.Account.DisplayName
			ev.Post.AuthorURL = status.Status.Reblog.Account.URL
			ev.Post.Actor = status.Status.Account.DisplayName
			ev.Post.ActorName = status.Status.Account.Username
		}

		mod.evchan <- ev
	}
}

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
			mod.handleStreamEvent(item)
		}
	}
}
