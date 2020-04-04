package main

import (
	"flag"
	"log"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/quickcontrols2"

	gap "github.com/muesli/go-app-paths"
	"github.com/muesli/telephant/accounts/mastodon"
)

var (
	debug = flag.Bool("debug", true, "Enable debug output")

	config     Config
	configFile string

	notificationModel    = NewMessageModel(nil)
	conversationModel    = NewMessageModel(nil)
	accountMessagesModel = NewMessageModel(nil)
	attachmentModel      = NewAttachmentModel(nil)
	paneModel            = NewPaneModel(nil)
)

// closePane closes a pane
func closePane(idx int64) {
	debugln("Closing pane", idx)
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
	accountBridge.SetAttachments(attachmentModel)
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
		pane.ID = "notifications"
		pane.Name = "Notifications"
		pane.Sticky = true
		pane.Default = true
		pane.Model = notificationModel
		paneModel.AddPane(pane)
	}
	{
		var pane = NewPane(nil)
		pane.ID = "home"
		pane.Name = "Messages"
		pane.Default = true
		pane.Model = postModel
		paneModel.AddPane(pane)
	}

	panes := tc.Panes()
	for _, p := range panes {
		if !p.Default {
			continue
		}

		model := NewMessageModel(nil)
		evchan := make(chan interface{})

		go handleEvents(evchan, model)
		p.Stream(evchan)

		var pane = NewPane(nil)
		pane.ID = p.ID
		pane.Name = p.Title
		pane.Default = p.Default
		pane.Model = model
		paneModel.AddPane(pane)
	}
	accountBridge.SetPanes(paneModel)

	evchan := make(chan interface{})
	go handleEvents(evchan, postModel)
	go tc.Run(evchan)
}

func debugln(s ...interface{}) {
	if *debug {
		log.Println(s...)
	}
}

func main() {
	flag.Parse()

	core.QCoreApplication_SetApplicationName("Telephant")
	core.QCoreApplication_SetOrganizationName("fribbledom.com")
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	ga := gui.NewQGuiApplication(len(os.Args), os.Args)
	ga.SetWindowIcon(gui.NewQIcon5(":/qml/images/telephant_logo.png"))
	setupQmlBridges()

	// load config
	scope := gap.NewScope(gap.User, "telephant")
	configDir, err := scope.ConfigPath("")
	if err != nil {
		panic(err)
	}
	os.MkdirAll(configDir, 0700)

	configFile, err = scope.ConfigPath("telephant.conf")
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
	configBridge.SetPositionX(config.PositionX)
	configBridge.SetPositionY(config.PositionY)
	configBridge.SetWidth(config.Width)
	configBridge.SetHeight(config.Height)

	if len(config.Account) > 0 {
		setupMastodon(config.Account[0])
	}
	runApp(config)

	// save config
	config.Theme = configBridge.Theme()
	config.Style = configBridge.Style()
	config.PositionX = configBridge.PositionX()
	config.PositionY = configBridge.PositionY()
	config.Width = configBridge.Width()
	config.Height = configBridge.Height()
	config.FirstRun = false
	SaveConfig(configFile, config)
}
