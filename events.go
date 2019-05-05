package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/muesli/chirp/accounts"
)

func linkify(in []byte) []byte {
	return []byte(fmt.Sprintf("<a style=\"text-decoration: none; color: orange;\" href=\"%s\">%s</a>", in, in))
}

func messageFromEvent(event accounts.MessageEvent) *Message {
	var p = NewMessage(nil)
	p.MessageID = event.Post.MessageID
	p.PostURL = event.Post.URL
	p.Name = event.Post.AuthorName
	p.Author = event.Post.Author
	p.AuthorURL = event.Post.AuthorURL
	p.Avatar = event.Post.Avatar
	p.Body = strings.TrimSpace(event.Post.Body)
	p.CreatedAt = event.Post.CreatedAt
	p.Reply = event.Reply
	p.ReplyToID = event.Post.ReplyToID
	p.ReplyToAuthor = event.Post.ReplyToAuthor
	p.Mention = event.Mention
	p.Like = event.Like
	p.Forward = event.Forward
	p.Actor = event.Post.Actor
	p.ActorName = event.Post.ActorName
	if len(event.Media) > 0 {
		p.Media = event.Media[0]
	}

	return p
}

// handleEvents handles incoming events and puts them into the right models
func handleEvents(eventsIn chan interface{}, messages *MessageModel, notifications *MessageModel) {
	for {
		ev, ok := <-eventsIn
		if !ok {
			log.Println()
			log.Println("Stopped event handler!")
			break
		}

		switch event := ev.(type) {
		case accounts.LoginEvent:
			{
				log.Println("Account login succeeded:", event.Username, event.Name, event.Avatar)
				accountBridge.SetUsername(event.Username)
				accountBridge.SetName(event.Name)
				accountBridge.SetAvatar(event.Avatar)
				accountBridge.SetProfileURL(event.ProfileURL)
				accountBridge.SetPosts(event.Posts)
				accountBridge.SetFollows(event.Follows)
				accountBridge.SetFollowers(event.Followers)
			}
		case accounts.MessageEvent:
			{
				// spw := &spew.ConfigState{Indent: "  ", DisableCapacities: true, DisablePointerAddresses: true}
				// log.Println("Message received:", spw.Sdump(event))

				p := messageFromEvent(event)

				// markup links
				/* re, err := regexp.Compile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-z]{2,16}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)
				if err != nil {
					log.Fatal("URL detection regexp does not compile: ", err)
				}
				p.Body = string(re.ReplaceAllFunc([]byte(p.Body), linkify)) */

				if event.Notification {
					notifications.AddMessage(p)
				} else {
					messages.AddMessage(p)
				}
			}
		}
	}
}
