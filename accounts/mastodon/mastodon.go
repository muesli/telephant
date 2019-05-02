// Package mastodon is a Mastodon account for Chirp.
package mastodon

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ChimeraCoder/anaconda"
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
	username     string
	password     string
	instance     string
	clientID     string
	clientSecret string

	client *mastodon.Client
	self   *mastodon.Account

	evchan  chan interface{}
	SigChan chan bool
}

// NewAccount returns a new Mastodon account.
func NewAccount(username, password, instance, clientID, clientSecret string) *Account {
	return &Account{
		username:     username,
		password:     password,
		instance:     instance,
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

// Run executes the account's event loop.
func (mod *Account) Run(eventChan chan interface{}) {
	mod.evchan = eventChan

	mod.client = mastodon.NewClient(&mastodon.Config{
		Server:       mod.instance,
		ClientID:     mod.clientID,
		ClientSecret: mod.clientSecret,
	})
	err := mod.client.Authenticate(context.Background(), mod.username, mod.password)
	if err != nil {
		panic(err)
	}

	mod.self, err = mod.client.GetAccountCurrentUser(context.Background())
	if err != nil {
		panic(err)
	}

	ev := accounts.LoginEvent{
		Username: mod.self.Username,
		Name:     mod.self.DisplayName,
		Avatar:   mod.self.Avatar,
	}
	mod.evchan <- ev

	// FIXME: retrieve initial feed
	mod.handleStream()
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
	return nil
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

func parseTweet(ents anaconda.Entities, ev *accounts.MessageEvent) {
	return

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
				Name:         "tweet",
				Notification: true,

				Post: accounts.Post{
					MessageID:  string(status.Notification.Status.ID),
					Body:       status.Notification.Status.Content,
					Author:     status.Notification.Account.DisplayName,
					AuthorName: status.Notification.Account.Username,
					AuthorURL:  status.Notification.Account.URL,
					Avatar:     status.Notification.Account.Avatar,
					CreatedAt:  time.Now(),
					URL:        status.Notification.Status.URL,
				},
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
			ev.Post.Author = status.Notification.Status.Account.DisplayName
			ev.Post.AuthorName = status.Notification.Status.Account.Username
			ev.Post.AuthorURL = status.Notification.Status.Account.URL
			// ev.Post.Avatar = status.Notification.Status.Account.Avatar
			ev.Post.Actor = status.Notification.Account.DisplayName
			ev.Post.ActorName = status.Notification.Account.Username

		case "favourite":
			ev.Like = true

			ev.Post.Author = status.Notification.Status.Account.DisplayName
			ev.Post.AuthorName = status.Notification.Status.Account.Username
			ev.Post.AuthorURL = status.Notification.Status.Account.URL
			// ev.Post.Avatar = status.Notification.Status.Account.Avatar
			ev.Post.Actor = status.Notification.Account.DisplayName
			ev.Post.ActorName = status.Notification.Account.Username

		default:
			fmt.Println("Unknown type:", status.Notification.Type)
			return
		}

		mod.evchan <- ev

	case *mastodon.UpdateEvent:
		ev := accounts.MessageEvent{
			Account: "mastodon",
			Name:    "tweet",
			Post: accounts.Post{
				MessageID:  string(status.Status.ID),
				Body:       status.Status.Content,
				Author:     status.Status.Account.DisplayName,
				AuthorName: status.Status.Account.Username,
				AuthorURL:  status.Status.Account.URL,
				Avatar:     status.Status.Account.Avatar,
				CreatedAt:  time.Now(),
				URL:        status.Status.URL,
			},
		}

		if status.Status.Reblog != nil {
			ev.Forward = true

			ev.Post.URL = status.Status.Reblog.URL
			ev.Post.Author = status.Status.Reblog.Account.DisplayName
			ev.Post.AuthorName = status.Status.Reblog.Account.Username
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
