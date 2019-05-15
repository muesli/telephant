package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"

	gap "github.com/muesli/go-app-paths"
	"github.com/muesli/telephant/accounts/mastodon"
)

var (
	config               Config
	notificationModel    = NewMessageModel(nil)
	conversationModel    = NewMessageModel(nil)
	accountMessagesModel = NewMessageModel(nil)
	paneModel            = NewPaneModel(nil)
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
		fmt.Println("Error registering app:", err)
		accountBridge.SetError(err.Error())
		return false
	}

	configBridge.SetAuthURL(authURI)
	configBridge.SetRedirectURL(redirectURI)

	fmt.Println("auth uri:", authURI)
	fmt.Println("redirect uri:", redirectURI)
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
		fmt.Println("Error authenticating with instance:", err)
		accountBridge.SetError(err.Error())
		return false
	}

	config.Account[0].Instance = instance
	config.Account[0].ClientID = clientID
	config.Account[0].ClientSecret = clientSecret
	config.Account[0].Token = token
	setupMastodon(config.Account[0])
	return true
}

func postLimitCount(body string) int {
	return tc.PostLimitCount(body)
}

// reply is used to post a new message
// if replyid is > 0, it's send as a reply
func reply(replyid string, message string) {
	var err error
	if replyid != "" {
		log.Println("Sending reply to:", replyid, message)
		err = tc.Reply(replyid, message)
	} else {
		log.Println("Posting:", message)
		err = tc.Post(message)
	}
	if err != nil {
		accountBridge.SetError(err.Error())
		log.Println("Error posting:", err)
	}
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

	fmt.Println("Found conversation posts:", len(messages))
	conversationModel.Clear()
	for _, m := range messages {
		p := messageFromEvent(m)
		conversationModel.AppendMessage(p)
	}
}

// loadAccount loads an entire profile
func loadAccount(id string) {
	log.Println("Loading account:", id)
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

	fmt.Println("Found account posts:", len(messages))
	accountMessagesModel.Clear()
	for _, m := range messages {
		p := messageFromEvent(m)
		accountMessagesModel.AppendMessage(p)
	}
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
	pane.Name = "Tag: #" + token
	pane.Model = model
	paneModel.AddPane(pane)
}

// closePane closes a pane
func closePane(idx int64) {
	fmt.Println("Closing pane", idx)
	paneModel.RemovePane(int(idx))
}

// runApp loads and executes the QML UI
func runApp(config Config) {
	var theme string
	switch config.Theme {
	case "System":
		theme = ""
	case "Light":
		theme = "Default"
	default:
		theme = config.Theme
	}
	if theme != "" {
		quickcontrols2.QQuickStyle_SetStyle(theme)
	}

	app := qml.NewQQmlApplicationEngine(nil)
	app.RootContext().SetContextProperty("uiBridge", uiBridge)
	app.RootContext().SetContextProperty("accountBridge", accountBridge)
	app.RootContext().SetContextProperty("profileBridge", profileBridge)
	app.RootContext().SetContextProperty("settings", configBridge)

	app.Load(core.NewQUrl3("qrc:/qml/telephant.qml", 0))
	gui.QGuiApplication_Exec()
}

// setupMastodon starts a new Mastodon client and sets up event handling & models for it
func setupMastodon(config Account) {
	tc = mastodon.NewAccount(config.Instance, config.Token, config.ClientID, config.ClientSecret)
	postModel := NewMessageModel(nil)

	accountBridge.SetUsername("Not connected...")
	accountBridge.SetNotifications(notificationModel)
	accountBridge.SetConversation(conversationModel)
	accountBridge.SetAccountMessages(accountMessagesModel)
	accountBridge.SetAvatar("qrc:/qml/images/telephant_logo.png")
	accountBridge.SetPosts(0)
	accountBridge.SetFollowCount(0)
	accountBridge.SetFollowerCount(0)
	accountBridge.SetPostSizeLimit(0)

	// Notifications model must the first model to be added
	// It will always be displayed right-most
	paneModel.clear()
	{
		var pane = NewPane(nil)
		pane.Name = "Notifications"
		pane.Sticky = true
		pane.Model = notificationModel
		paneModel.AddPane(pane)
	}
	{
		var pane = NewPane(nil)
		pane.Name = "Messages"
		pane.Model = postModel
		paneModel.AddPane(pane)
	}

	panes := tc.Panes()
	for _, p := range panes {
		model := NewMessageModel(nil)
		evchan := make(chan interface{})

		go handleEvents(evchan, model)
		p.Stream(evchan)

		var pane = NewPane(nil)
		pane.Name = p.Title
		pane.Model = model
		paneModel.AddPane(pane)
	}
	accountBridge.SetPanes(paneModel)

	evchan := make(chan interface{})
	go handleEvents(evchan, postModel)
	go tc.Run(evchan)
}

func main() {
	core.QCoreApplication_SetApplicationName("Telephant")
	core.QCoreApplication_SetOrganizationName("fribbledom.com")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	ga := gui.NewQGuiApplication(len(os.Args), os.Args)
	ga.SetWindowIcon(gui.NewQIcon5(":/qml/images/telephant_logo.png"))
	setupQmlBridges()

	// load config
	scope := gap.NewScope(gap.User, "fribbledom.com", "telephant")
	configDir, err := scope.ConfigPath("")
	if err != nil {
		panic(err)
	}
	os.MkdirAll(configDir, 0700)

	configFile, err := scope.ConfigPath("telephant.conf")
	if err != nil {
		panic(err)
	}
	config = LoadConfig(configFile)
	if config.Theme == "" {
		config.Theme = "Material"
	}
	if config.Style == "" {
		config.Style = "Dark"
	}
	configBridge.SetTheme(config.Theme)
	configBridge.SetStyle(config.Style)
	configBridge.SetFirstRun(config.FirstRun)

	setupMastodon(config.Account[0])
	runApp(config)

	// save config
	config.Theme = configBridge.Theme()
	config.Style = configBridge.Style()
	config.FirstRun = false
	SaveConfig(configFile, config)
}
