// Package accounts is telephant's account "plugin" system.
package accounts

import "time"

// Post represents a message.
type Post struct {
	MessageID     string
	Body          string
	Sensitive     bool
	Warning       string
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
	PostID        string
	URL           string
	CreatedAt     time.Time
	Liked         bool
	Shared        bool
	RepliesCount  int64
	LikesCount    int64
	SharesCount   int64
}

// Follow describes an incoming follow event.
type Follow struct {
	Account    string
	Name       string
	Avatar     string
	ProfileURL string
	ProfileID  string
	Following  bool
	FollowedBy bool
}

// Media describes a media item.
type Media struct {
	ID      string
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
	Notify       bool
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

type DeleteEvent struct {
	ID string
}

// ErrorEvent describes an error that occurred.
type ErrorEvent struct {
	Message  string
	Internal bool
}
