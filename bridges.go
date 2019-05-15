package main

import (
	"github.com/muesli/telephant/accounts/mastodon"
	"github.com/therecipe/qt/core"
)

// UIBridge lets us trigger Go methods from QML
type UIBridge struct {
	core.QObject

	_ func(instance string) bool                 `slot:"connectButton"`
	_ func(code string, redirectURI string) bool `slot:"authButton"`

	_ func(replyid string, message string) `slot:"postButton"`
	_ func(id string)                      `slot:"deleteButton"`
	_ func(id string)                      `slot:"shareButton"`
	_ func(id string)                      `slot:"unshareButton"`
	_ func(id string)                      `slot:"likeButton"`
	_ func(id string)                      `slot:"unlikeButton"`
	_ func(id string, follow bool)         `slot:"followButton"`
	_ func(id string)                      `slot:"loadConversation"`
	_ func(id string)                      `slot:"loadAccount"`
	_ func(token string)                   `slot:"tag"`

	_ func(idx int64) `slot:"closePane"`

	_ func(object *core.QObject) `slot:"registerToGo"`
	_ func(objectName string)    `slot:"deregisterToGo"`
}

// AccountBridge makes an account plugin available in QML
type AccountBridge struct {
	core.QObject

	_ string `property:"username"`
	_ string `property:"name"`
	_ string `property:"avatar"`
	_ string `property:"profileURL"`
	_ string `property:"profileID"`
	_ int64  `property:"posts"`
	_ int64  `property:"followCount"`
	_ int64  `property:"followerCount"`
	_ int64  `property:"postSizeLimit"`

	_ string `property:"error"`

	_ *core.QAbstractListModel `property:"panes"`
	_ *core.QAbstractListModel `property:"messages"`
	_ *core.QAbstractListModel `property:"notifications"`
	_ *core.QAbstractListModel `property:"conversation"`
	_ *core.QAbstractListModel `property:"accountMessages"`
}

// ProfileBridge allows QML to access profile data
type ProfileBridge struct {
	core.QObject

	_ string `property:"username"`
	_ string `property:"name"`
	_ string `property:"avatar"`
	_ string `property:"profileURL"`
	_ string `property:"profileID"`
	_ int64  `property:"posts"`
	_ int64  `property:"followCount"`
	_ int64  `property:"followerCount"`
	_ bool   `property:"following"`
	_ bool   `property:"followedBy"`
}

// ConfigBridge allows QML to access the app's config
type ConfigBridge struct {
	core.QObject

	_ bool   `property:"firstRun"`
	_ string `property:"authURL"`
	_ string `property:"redirectURL"`
	_ string `property:"theme"`
	_ string `property:"style"`
}

var (
	// qmlObjects    = make(map[string]*core.QObject)
	uiBridge      *UIBridge
	accountBridge *AccountBridge
	configBridge  *ConfigBridge
	profileBridge *ProfileBridge
	tc            *mastodon.Account
)

// setupQmlBridges initializes the QML bridges
func setupQmlBridges() {
	configBridge = NewConfigBridge(nil)

	accountBridge = NewAccountBridge(nil)
	accountBridge.SetUsername("Telephant!")

	uiBridge = NewUIBridge(nil)
	uiBridge.ConnectConnectButton(connectToInstance)
	uiBridge.ConnectAuthButton(authInstance)
	uiBridge.ConnectPostButton(reply)
	uiBridge.ConnectDeleteButton(deletePost)
	uiBridge.ConnectShareButton(share)
	uiBridge.ConnectUnshareButton(unshare)
	uiBridge.ConnectLikeButton(like)
	uiBridge.ConnectUnlikeButton(unlike)
	uiBridge.ConnectFollowButton(follow)
	uiBridge.ConnectLoadConversation(loadConversation)
	uiBridge.ConnectLoadAccount(loadAccount)
	uiBridge.ConnectTag(tag)
	uiBridge.ConnectClosePane(closePane)

	profileBridge = NewProfileBridge(nil)

	/*	uiBridge.ConnectRegisterToGo(func(object *core.QObject) {
			qmlObjects[object.ObjectName()] = object
		})
		uiBridge.ConnectDeregisterToGo(func(objectName string) {
			qmlObjects[objectName] = nil
		}) */
}
