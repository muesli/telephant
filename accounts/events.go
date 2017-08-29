// Package accounts is chirp's account "plugin" system.
package accounts

import "time"

// Post represents a message.
type Post struct {
	MessageID     int64
	Body          string
	Author        string
	AuthorName    string
	Actor         string
	ActorName     string
	ReplyToID     int64
	ReplyToAuthor string
	Avatar        string
	URL           string
	CreatedAt     time.Time
}

// MessageEvent describes an incoming message event.
type MessageEvent struct {
	Account      string
	Name         string
	Reply        bool
	Forward      bool
	Mention      bool
	Like         bool
	Notification bool
	Post         Post
}

// LoginEvent describes a login event.
type LoginEvent struct {
	Username string
	Name     string
	Avatar   string
}
