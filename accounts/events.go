// Package accounts is telephant's account "plugin" system.
package accounts

import "time"

// Post represents a message.
type Post struct {
	MessageID     string
	Body          string
	Author        string
	AuthorURL     string
	AuthorName    string
	AuthorID      string
	Actor         string
	ActorName     string
	ActorID       string
	ReplyToID     string
	ReplyToAuthor string
	Avatar        string
	URL           string
	CreatedAt     time.Time
	Liked         bool
	Shared        bool
}

type Follow struct {
	Account    string
	Name       string
	Avatar     string
	ProfileURL string
	ProfileID  string
	Following  bool
	FollowedBy bool
}

type Media struct {
	Preview string
	URL     string
}

// MessageEvent describes an incoming message event.
type MessageEvent struct {
	Account      string
	Name         string
	Reply        bool
	Forward      bool
	Mention      bool
	Like         bool
	Followed     bool
	Notification bool
	Post         Post
	Follow       Follow
	Media        []Media
}

// ProfileEvent describes a profile event.
type ProfileEvent struct {
	Username      string
	Name          string
	Avatar        string
	ProfileURL    string
	ProfileID     string
	Posts         int64
	FollowCount   int64
	FollowerCount int64
	Following     bool
	FollowedBy    bool
}

// LoginEvent describes a login event.
type LoginEvent struct {
	Username      string
	Name          string
	Avatar        string
	ProfileURL    string
	ProfileID     string
	Posts         int64
	FollowCount   int64
	FollowerCount int64
	PostSizeLimit int64
}

// ErrorEvent describes an error that occurred.
type ErrorEvent struct {
	Message  string
	Internal bool
}
