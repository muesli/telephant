package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/muesli/telephant/accounts"
	"github.com/tachiniererin/bananasplit"
)

func addEmojiFont(body string) string {
	runeRange := map[string][]bananasplit.RuneRange{"emoji": bananasplit.EmojiRange}
	bodyParts := bananasplit.SplitByRanges(body, runeRange)
	var mon strings.Builder
	for _, part := range bodyParts {
		if part.Type == "emoji" {
			mon.WriteString(fmt.Sprintf(`<font face="%s">%s</font>`, config.EmojiFont, part.Text))
		} else {
			mon.WriteString(part.Text)
		}
	}

	return mon.String()
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
				accountBridge.SetName(addEmojiFont(event.Name))
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
		case accounts.DeleteEvent:
			{
				deleteMessage(event.ID)
			}
		case accounts.MessageEvent:
			{
				// spw := &spew.ConfigState{Indent: "  ", DisableCapacities: true, DisablePointerAddresses: true}
				// log.Println("Message received:", spw.Sdump(event))

				p := messageFromEvent(event)
				p.Name = addEmojiFont(p.Name)
				p.ActorName = addEmojiFont(p.ActorName)
				p.Body = addEmojiFont(p.Body)

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
