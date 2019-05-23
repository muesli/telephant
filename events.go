package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/muesli/telephant/accounts"
)

// messageFromEvent creates a new Message object from an incoming MessageEvent.
func messageFromEvent(event accounts.MessageEvent) *Message {
	var p = NewMessage(nil)
	p.MessageID = event.Post.MessageID
	p.PostURL = event.Post.URL
	p.Name = event.Post.AuthorName
	p.Author = event.Post.Author
	p.AuthorURL = event.Post.AuthorURL
	p.AuthorID = event.Post.AuthorID
	p.Avatar = event.Post.Avatar
	p.Body = strings.TrimSpace(event.Post.Body)
	p.CreatedAt = event.Post.CreatedAt
	p.Reply = event.Reply
	p.ReplyToID = event.Post.ReplyToID
	p.ReplyToAuthor = event.Post.ReplyToAuthor
	p.Forward = event.Forward
	p.Mention = event.Mention
	p.Like = event.Like
	p.Followed = event.Followed
	p.Actor = event.Post.Actor
	p.ActorName = event.Post.ActorName
	p.ActorID = event.Post.ActorID
	p.Liked = event.Post.Liked
	p.Shared = event.Post.Shared
	if len(event.Media) > 0 {
		for _, v := range event.Media {
			p.MediaPreview = append(p.MediaPreview, v.Preview)
			p.MediaURL = append(p.MediaURL, v.URL)
		}
	}

	if p.Followed {
		p.Actor = event.Follow.Account
		p.ActorName = event.Follow.Name
		p.Avatar = event.Follow.Avatar
		p.AuthorURL = event.Follow.ProfileURL
		p.AuthorID = event.Follow.ProfileID
		p.Following = event.Follow.Following
		p.FollowedBy = event.Follow.FollowedBy
	}

	return p
}

// handleEvents handles incoming events and puts them into the right models
func handleEvents(eventsIn chan interface{}, messages *MessageModel) {
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
				accountBridge.SetProfileID(event.ProfileID)
				accountBridge.SetPosts(event.Posts)
				accountBridge.SetFollowCount(event.FollowCount)
				accountBridge.SetFollowerCount(event.FollowerCount)
				accountBridge.SetPostSizeLimit(event.PostSizeLimit)
			}
		case accounts.ErrorEvent:
			{
				log.Println("Error:", event.Message)
				accountBridge.SetError(event.Message)
			}
		case accounts.Media:
			{
				log.Printf("Added attachment: %+v\n", event)
				var p = NewAttachment(nil)
				p.ID = event.ID
				p.Preview = event.Preview
				attachmentModel.AddAttachment(p)
			}
		case accounts.MessageEvent:
			{
				// spw := &spew.ConfigState{Indent: "  ", DisableCapacities: true, DisablePointerAddresses: true}
				// log.Println("Message received:", spw.Sdump(event))

				p := messageFromEvent(event)

				// markup links
				/*
					re, err := regexp.Compile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-z]{2,16}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)`)
					if err != nil {
						log.Fatal("URL detection regexp does not compile: ", err)
					}
					p.Body = string(re.ReplaceAllFunc([]byte(p.Body), linkify))
				*/

				if event.Notification {
					notificationModel.AddMessage(p)

					if event.Notify {
						title := "Telephant"
						body := p.Body
						if p.Mention {
							body = fmt.Sprintf("%s mentioned you", p.Name)
						}
						if p.Followed {
							body = fmt.Sprintf("%s followed you", p.ActorName)
						}
						if p.Like {
							body = fmt.Sprintf("%s liked your post", p.ActorName)
						}
						if p.Forward {
							body = fmt.Sprintf("%s shared your post", p.ActorName)
						}

						/*
							if body == "" {
								spw := &spew.ConfigState{Indent: "  ", DisableCapacities: true, DisablePointerAddresses: true}
								log.Println("Unknown notification received:", spw.Sdump(event))
							}
						*/

						if body != "" {
							notify(title, body)
						}
					}
				} else {
					messages.AddMessage(p)
				}
			}
		}
	}
}

// linkify is a helper function to wrap an HTML anchor around detected links.
func linkify(in []byte) []byte {
	return []byte(fmt.Sprintf("<a style=\"text-decoration: none;\" href=\"%s\">%s</a>", in, in))
}
