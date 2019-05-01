package main

import (
	"github.com/muesli/chirp/accounts/mastodon"
	"github.com/therecipe/qt/core"
)

// UIBridge lets us trigger Go methods from QML
type UIBridge struct {
	core.QObject

	_ func(replyid string, message string) `slot:"postButton"`
	_ func(id string)                      `slot:"retweetButton"`
	_ func(id string)                      `slot:"likeButton"`

	_ func(object *core.QObject) `slot:"registerToGo"`
	_ func(objectName string)    `slot:"deregisterToGo"`
}

// AccountBridge makes an account plugin available in QML
type AccountBridge struct {
	core.QObject

	_ string `property:"username"`
	_ string `property:"avatar"`

	_ *core.QAbstractListModel `property:"messages"`
	_ *core.QAbstractListModel `property:"notifications"`
}

// ConfigBridge allows QML to access the app's config
type ConfigBridge struct {
	core.QObject

	_ string `property:"style"`
}

var (
	// qmlObjects    = make(map[string]*core.QObject)
	uiBridge      *UIBridge
	accountBridge *AccountBridge
	configBridge  *ConfigBridge
	tc            *mastodon.Account
)

// setupQmlBridges initializes the QML bridges
func setupQmlBridges() {
	configBridge = NewConfigBridge(nil)

	accountBridge = NewAccountBridge(nil)
	accountBridge.SetUsername("Chirp!")

	uiBridge = NewUIBridge(nil)
	uiBridge.ConnectPostButton(reply)
	uiBridge.ConnectRetweetButton(retweet)
	uiBridge.ConnectLikeButton(like)

	/*	uiBridge.ConnectRegisterToGo(func(object *core.QObject) {
			qmlObjects[object.ObjectName()] = object
		})
		uiBridge.ConnectDeregisterToGo(func(objectName string) {
			qmlObjects[objectName] = nil
		}) */
}
