package main

import (
	"log"
	"os"
	"strconv"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"

	"github.com/muesli/chirp/accounts/twitter"
)

// reply is used to post a new tweet
// if replyid is > 0, it's send as a reply
func reply(replyid string, message string) {
	var err error
	iid, _ := strconv.ParseInt(replyid, 10, 64)
	if iid > 0 {
		log.Println("Sending reply to:", iid, message)
		err = tc.Reply(iid, message)
	} else {
		log.Println("Sending tweet:", message)
		err = tc.Tweet(message)
	}
	if err != nil {
		log.Println("Error posting to Twitter:", err)
	}
}

// retweet a message
func retweet(id string) {
	iid, _ := strconv.ParseInt(id, 10, 64)
	log.Println("Retweeting:", iid)
	err := tc.Retweet(iid)
	log.Println("Error posting to Twitter:", err)
}

// like a message
func like(id string) {
	iid, _ := strconv.ParseInt(id, 10, 64)
	log.Println("Liking:", iid)
	err := tc.Like(iid)
	log.Println("Error posting to Twitter:", err)
}

// runApp loads and executes the QML UI
func runApp(config Config) {
	quickcontrols2.QQuickStyle_SetStyle(config.Style)

	app := qml.NewQQmlApplicationEngine(nil)
	app.RootContext().SetContextProperty("uiBridge", uiBridge)
	app.RootContext().SetContextProperty("accountBridge", accountBridge)
	app.RootContext().SetContextProperty("settings", configBridge)

	app.Load(core.NewQUrl3("qrc:/qml/chirp.qml", 0))
	gui.QGuiApplication_Exec()
}

// setupTwitter starts a new Twitter client and sets up event handling & models for it
func setupTwitter(config Account) {
	tc = twitter.NewAccount(config.ConsumerKey, config.ConsumerSecret, config.AccessToken, config.AccessTokenSecret)
	tweetModel := NewMessageModel(nil)
	notificationModel := NewMessageModel(nil)

	accountBridge.SetUsername("Logging in...")
	accountBridge.SetMessages(tweetModel)
	accountBridge.SetNotifications(notificationModel)

	evchan := make(chan interface{})
	go handleEvents(evchan, tweetModel, notificationModel)
	go func() {
		tc.Run(evchan)
	}()
}

func main() {
	core.QCoreApplication_SetApplicationName("Chirp")
	core.QCoreApplication_SetOrganizationName("chris.de")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	gui.NewQGuiApplication(len(os.Args), os.Args)

	setupQmlBridges()

	// load config
	config := LoadConfig()
	if config.Style == "" {
		config.Style = "Material"
	}
	configBridge.SetStyle(config.Style)

	setupTwitter(config.Account[0])
	runApp(config)

	// save config
	config.Style = configBridge.Style()
	SaveConfig(config)
}
