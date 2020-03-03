package main

import (
	"log"
	"net/url"
	"strings"

	"github.com/muesli/telephant/accounts/mastodon"
)

// connectToInstance registers an app with the instance and retrieves an
// authentication URI.
func connectToInstance(instance string) bool {
	var authURI string
	var redirectURI string
	var err error
	instance = addHTTPPrefixIfNeeded(instance)
	tc, authURI, redirectURI, err = mastodon.RegisterAccount(instance)
	if err != nil {
		log.Println("Error registering app:", err)
		accountBridge.SetError(err.Error())
		return false
	}

	configBridge.SetAuthURL(authURI)
	configBridge.SetRedirectURL(redirectURI)

	log.Println("auth uri:", authURI)
	log.Println("redirect uri:", redirectURI)
	return true
}

// addHTTPPrefixIfNeeded adds "https://" to an instance URL where it's missing.
func addHTTPPrefixIfNeeded(instance string) string {
	if !strings.HasPrefix(instance, "http://") && !strings.HasPrefix(instance, "https://") {
		return "https://" + instance
	}

	return instance
}

// authInstance authenticates a user via OAuth and retrieves an accesstoken
// we'll use for future logins.
func authInstance(code, redirectURI string) bool {
	instance, token, clientID, clientSecret, err := tc.Authenticate(code, redirectURI)
	if err != nil {
		log.Println("Error authenticating with instance:", err)
		accountBridge.SetError(err.Error())
		return false
	}

	config.Account[0].Instance = instance
	config.Account[0].ClientID = clientID
	config.Account[0].ClientSecret = clientSecret
	config.Account[0].Token = token
	config.FirstRun = false
	SaveConfig(configFile, config)

	setupMastodon(config.Account[0])
	return true
}

func postLimitCount(body string) int {
	return tc.PostLimitCount(body)
}

// reply is used to post a new message
// if replyid is > 0, it's send as a reply
func reply(replyid string, message string) {
	var attachments []string
	for _, v := range attachmentModel.Attachments() {
		attachments = append(attachments, v.ID)
	}

	var err error
	if replyid != "" {
		log.Println("Sending reply to:", replyid, attachments, message)
		err = tc.Reply(replyid, message, attachments)
	} else {
		log.Println("Posting:", attachments, message)
		err = tc.Post(message, attachments)
	}
	if err != nil {
		accountBridge.SetError(err.Error())
		log.Println("Error posting:", err)
	}
}

func uploadAttachment(pathurl string) {
	u, _ := url.ParseRequestURI(pathurl)
	log.Println("Uploding:", u.Path)
	tc.UploadAttachment(u.Path)
}

// deletePost deletes a post
func deletePost(id string) {
	log.Println("Deleting:", id)
	if err := tc.DeletePost(id); err != nil {
		accountBridge.SetError(err.Error())
		log.Println("Error deleting:", err)
	}
}

// share a post
func share(id string) {
	log.Println("Sharing:", id)
	if err := tc.Share(id); err != nil {
		accountBridge.SetError(err.Error())
		log.Println("Error sharing:", err)
	}
}

// unshare a post
func unshare(id string) {
	log.Println("Unsharing:", id)
	if err := tc.Unshare(id); err != nil {
		accountBridge.SetError(err.Error())
		log.Println("Error unsharing:", err)
	}
}

// like a post
func like(id string) {
	log.Println("Liking:", id)
	if err := tc.Like(id); err != nil {
		accountBridge.SetError(err.Error())
		log.Println("Error liking:", err)
	}
}

// unlike a post
func unlike(id string) {
	log.Println("Unliking:", id)
	if err := tc.Unlike(id); err != nil {
		accountBridge.SetError(err.Error())
		log.Println("Error unliking:", err)
	}
}

// follow changes the relationship to another user
func follow(id string, follow bool) {
	if follow {
		log.Println("Following:", id)
		if err := tc.Follow(id); err != nil {
			accountBridge.SetError(err.Error())
			log.Println("Error following user:", err)
			return
		}
	} else {
		log.Println("Unfollowing:", id)
		if err := tc.Unfollow(id); err != nil {
			accountBridge.SetError(err.Error())
			log.Println("Error unfollowing user:", err)
			return
		}
	}

	profileBridge.SetFollowing(follow)
}

// loadConversation loads a message thread
func loadConversation(id string) {
	log.Println("Loading conversation:", id)
	messages, err := tc.LoadConversation(id)
	if err != nil {
		log.Println("Error loading conversation:", err)
		return
	}

	debugln("Found conversation posts:", len(messages))
	conversationModel.Clear()
	for _, m := range messages {
		p := messageFromEvent(m)
		conversationModel.AppendMessage(p)
	}
}

// loadAccount loads an entire profile
func loadAccount(id string) {
	debugln("Loading account:", id)
	profile, messages, err := tc.LoadAccount(id)
	if err != nil {
		log.Println("Error loading account:", err)
		return
	}

	profileBridge.SetUsername(profile.Username)
	profileBridge.SetName(profile.Name)
	profileBridge.SetAvatar(profile.Avatar)
	profileBridge.SetProfileURL(profile.ProfileURL)
	profileBridge.SetProfileID(profile.ProfileID)
	profileBridge.SetPosts(profile.Posts)
	profileBridge.SetFollowCount(profile.FollowCount)
	profileBridge.SetFollowerCount(profile.FollowerCount)
	profileBridge.SetFollowing(profile.Following)
	profileBridge.SetFollowedBy(profile.FollowedBy)

	debugln("Found account posts:", len(messages))
	accountMessagesModel.Clear()
	for _, m := range messages {
		p := messageFromEvent(m)
		accountMessagesModel.AppendMessage(p)
	}
}

// search
func search(token string) {
	model := NewMessageModel(nil)
	evchan := make(chan interface{})
	go handleEvents(evchan, model)

	log.Println("Search:", token)
	if err := tc.Search(token, evchan); err != nil {
		accountBridge.SetError(err.Error())
		log.Println("Error retrieving search:", err)
		return
	}

	var pane = NewPane(nil)
	pane.ID = "search_" + token
	pane.Name = "Search: " + token
	pane.Model = model
	paneModel.AddPane(pane)
}

// tag
func tag(token string) {
	model := NewMessageModel(nil)
	evchan := make(chan interface{})
	go handleEvents(evchan, model)

	log.Println("Hashtag:", token)
	if err := tc.Tag(token, evchan); err != nil {
		accountBridge.SetError(err.Error())
		log.Println("Error retrieving hashtag:", err)
		return
	}

	var pane = NewPane(nil)
	pane.ID = "tag_" + token
	pane.Name = "Tag: #" + token
	pane.Model = model
	paneModel.AddPane(pane)
}
