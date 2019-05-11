// Package accounts is telephant's account "plugin" system.
package accounts

type Account interface {
	Logo() string
	Authenticate(code, redirectURI string) (string, string, string, string, error)
	Run(eventChan chan interface{}) error

	Post(message string) error
	Reply(replyid string, message string) error
	Share(id string) error
	Unshare(id string) error
	Like(id string) error
	Unlike(id string) error
	Follow(id string) error
	Unfollow(id string) error

	LoadConversation(id string) ([]MessageEvent, error)
	LoadAccount(id string) (ProfileEvent, []MessageEvent, error)
}
