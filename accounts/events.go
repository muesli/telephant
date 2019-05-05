// Package accounts is chirp's account "plugin" system.
package accounts

import "time"

// Post represents a message.
type Post struct {
	MessageID     string
	Body          string
	Author        string
	AuthorURL     string
	AuthorName    string
	Actor         string
	ActorName     string
	ReplyToID     string
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
	Media        []string
}

// LoginEvent describes a login event.
type LoginEvent struct {
	Username   string
	Name       string
	Avatar     string
	ProfileURL string
	Posts      int64
	Follows    int64
	Followers  int64
}
