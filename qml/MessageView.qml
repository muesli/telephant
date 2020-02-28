import QtQuick 2.5
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3
import QtGraphicalEffects 1.0

import "componentCreator.js" as ComponentCreator

ColumnLayout {
    property bool fadeMedia
    property bool showActionButtons: true
    property bool showSensitiveContent: false
    property var message: model

    property bool following: message.following
    property bool liked: message.liked
    property bool shared: message.shared

    clip: true

    RowLayout {
        visible: message.forward && !message.like
        Item {
            width: 32
        }
        Image {
            smooth: true
            source: "images/share.png"
            sourceSize.height: 14
            opacity: 0.5
        }
        Label {
            font.pointSize: 10
            text: qsTr("%1 shared").arg(message.actorname)
            opacity: (accountBridge.username == message.author && (message.like || message.forward)) ? 0.8 : 0.3

            MouseArea {
                anchors.fill: parent
                cursorShape: Qt.PointingHandCursor
                onClicked: {
                    uiBridge.loadAccount(message.actorid)
                    ComponentCreator.createAccountPopup(this).open();
                }
            }
        }
    }
    RowLayout {
        visible: message.like
        Item {
            width: 32
        }
        Image {
            smooth: true
            source: "images/like.png"
            sourceSize.height: 14
            opacity: 0.5
        }
        Label {
            font.pointSize: 10
            text: qsTr("%1 liked").arg(message.actorname)
            opacity: (accountBridge.username == message.author && (message.like || message.forward)) ? 0.8 : 0.3

            MouseArea {
                anchors.fill: parent
                cursorShape: Qt.PointingHandCursor
                onClicked: {
                    uiBridge.loadAccount(message.actorid)
                    ComponentCreator.createAccountPopup(this).open();
                }
            }
        }
    }

    RowLayout {
        spacing: 8
        Layout.fillWidth: true

        ImageButton {
            id: image
            Layout.alignment: Qt.AlignTop
            source: message.avatar
            sourceSize.width: 48
            width: 48
            fillMode: Image.PreserveAspectCrop
            roundness: 4
            opacity: 1.0

            onClicked: function() {
                uiBridge.loadAccount(message.authorid)
                ComponentCreator.createAccountPopup(this).open();
            }
        }
        RowLayout {
            visible: message.followed
            Layout.fillWidth: true
            spacing: 4

            ColumnLayout {
                Layout.fillWidth: true
                Label {
                    font.pointSize: 11
                    font.bold: true
                    text: qsTr("%1 followed you").arg(message.actorname)
                    textFormat: Text.PlainText
                    Layout.fillWidth: true
                    elide: Text.ElideRight

                    MouseArea {
                        anchors.fill: parent
                        cursorShape: Qt.PointingHandCursor
                        onClicked: {
                            uiBridge.loadAccount(message.authorid)
                            ComponentCreator.createAccountPopup(this).open();
                        }
                    }
                }
                Label {
                    font.pointSize: 11
                    text: message.actor
                    textFormat: Text.PlainText
                    Layout.fillWidth: true
                    elide: Text.ElideRight

                    MouseArea {
                        anchors.fill: parent
                        cursorShape: Qt.PointingHandCursor
                        onClicked: {
                            uiBridge.loadAccount(message.authorid)
                            ComponentCreator.createAccountPopup(this).open();
                        }
                    }
                }
            }
            Button {
                Layout.alignment: Qt.AlignVCenter | Qt.AlignRight
                highlighted: true
                text: following ? qsTr("Unfollow") : qsTr("Follow")

                onClicked: {
                    uiBridge.followButton(message.authorid, !following)
                    following = !following
                }
            }
        }
        ColumnLayout {
            visible: !message.followed
            Layout.fillWidth: true
            spacing: 4

            RowLayout {
                width: parent.width
                Label {
                    id: namelabel
                    font.pointSize: 11
                    font.bold: true
                    text: message.name
                    textFormat: Text.PlainText
                    elide: Text.ElideRight
                    opacity: (accountBridge.username == message.author && (message.like || message.forward)) ? 0.4 : 1.0
                    Layout.fillWidth: true
                    Layout.maximumWidth: implicitWidth + 1

                    MouseArea {
                        anchors.fill: parent
                        cursorShape: Qt.PointingHandCursor
                        onClicked: {
                            uiBridge.loadAccount(message.authorid)
                            ComponentCreator.createAccountPopup(this).open();
                        }
                    }
                }
                Label {
                    // anchors.bottom: parent.bottom
                    font.pointSize: 9
                    opacity: 0.4
                    text: "@" + message.author
                    textFormat: Text.PlainText
                    elide: Text.ElideRight
                    Layout.fillWidth: true
                    Layout.maximumWidth: implicitWidth + 1

                    MouseArea {
                        anchors.fill: parent
                        cursorShape: Qt.PointingHandCursor
                        onClicked: {
                            uiBridge.loadAccount(message.authorid)
                            ComponentCreator.createAccountPopup(this).open();
                        }
                    }
                }
                Item {
                    // spacer item
                    Layout.fillWidth: true
                }
                Label {
                    font.pointSize: 9
                    opacity: 0.4
                    text: message.createdat
                    width: implicitWidth + 1
                    horizontalAlignment: Text.AlignRight

                    MouseArea {
                        anchors.fill: parent
                        cursorShape: Qt.PointingHandCursor
                        onClicked: {
                            uiBridge.loadConversation(message.postid)
                            ComponentCreator.createConversationPopup(this).open();
                        }
                    }
                }
            }
            ColumnLayout {
                Layout.fillWidth: true
                visible: !message.followed
                // width: parent.width
                // anchors.bottom: parent.bottom
                spacing: 4
                Label {
                    visible: message.sensitive && message.warning.length > 0
                    text: message.warning
                    font.pointSize: 11
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                }
                Button {
                    visible: message.sensitive && !showSensitiveContent
                    Layout.alignment: Qt.AlignHCenter
                    highlighted: true
                    text: qsTr("Show Content")

                    onClicked: {
                        showSensitiveContent = !showSensitiveContent
                    }
                }

                MessageText {
                    visible: message.body.length > 0 && (!message.sensitive || showSensitiveContent)
                    text: "<style>a:link { visibility: hidden; text-decoration: none; color: " + Material.accent + "; }</style>" + message.body
                    textFormat: Text.RichText
                    font.pointSize: 11
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                    opacity: (accountBridge.username == message.author && (message.like || message.forward)) ? 0.4 : 1.0
                    color: "white"

                    onLinkActivated: function(link) {
                        if (link.startsWith("telephant://")) {
                            var us = link.substr(12, link.length).split("/")

                            if (us[1] == "user") {
                                uiBridge.loadAccount(us[us.length-1])
                                ComponentCreator.createAccountPopup(this).open();
                            }
                            if (us[1] == "tag") {
                                uiBridge.tag(us[us.length-1])
                            }
                        } else
                            Qt.openUrlExternally(link)
                    }

                    onClicked: function() {
                            uiBridge.loadConversation(message.messageid)
                            ComponentCreator.createConversationPopup(this).open();
                    }
                }

                Flow {
                    id: flowgrid
                    visible: message.mediapreview.length > 0 && (!message.sensitive || showSensitiveContent)
                    Layout.fillWidth: true
                    Layout.topMargin: 4

                    property int cols: message.mediapreview.length >= 2 ? Math.min(message.mediapreview.length, width / 140) : 1
                    spacing: 4

                    Repeater {
                        model: message.mediapreview

                        ImageButton {
                            source: modelData
                            height: Math.min(sourceSize.height, flowgrid.width / 2)
                            width: Math.min(sourceSize.width, flowgrid.width / flowgrid.cols - flowgrid.spacing)
                            fillMode: Image.PreserveAspectCrop
                            verticalAlignment: Image.AlignVCenter
                            horizontalAlignment: Image.AlignHCenter
                            autoTransform: true
                            opacity: fadeMedia ? 0.2 : 1.0
                            roundness: 4
                            animationDuration: 200

                            onClicked: function() {
                                ComponentCreator.createMediaPopup(this, message.mediaurl[index]).open();
                                // Qt.openUrlExternally(message.mediaurl[index])
                            }
                        }
                    }
                }

                RowLayout {
                    Layout.fillWidth: true
                    RowLayout {
                        visible: message.reply
                        Label {
                            font.pointSize: 10
                            text: qsTr("Replying to %1").arg(
                                      "@" + message.replytoauthor)
                            opacity: 0.4

                            MouseArea {
                                anchors.fill: parent
                                cursorShape: Qt.PointingHandCursor
                                onClicked: {
                                    Qt.openUrlExternally(message.posturl)
                                }
                            }
                        }
                    }

                    RowLayout {
                        width: parent.width
                        visible: showActionButtons && !(accountBridge.username == message.author && (message.like || message.forward))
                        Layout.topMargin: 4

                        Item {
                            // spacer item
                            Layout.fillWidth: true
                        }
                        ImageButton {
                            source: "images/reply.png"
                            animationDuration: 200
                            sourceSize.height: 16
                            onClicked: function () {
                                ComponentCreator.createMessagePopup(this, message).open();
                            }
                        }
                        Label {
                            text: message.repliescount
                            font.pointSize: 9
                            opacity: 0.4
                        }
                        Label {
                            text: "·"
                            font.pointSize: 9
                            opacity: 0.4
                        }
                        ImageButton {
                            source: "images/share.png"
                            animationDuration: 200
                            sourceSize.height: 16
                            opacity: shared ? 1.0 : 0.3
                            onClicked: function () {
                                if (shared) {
                                    uiBridge.unshareButton(message.messageid)
                                    shared = false
                                } else {
                                    ComponentCreator.createSharePopup(this, message).open();
                                }
                            }
                        }
                        Label {
                            text: message.sharescount
                            font.pointSize: 9
                            opacity: 0.4
                        }
                        Label {
                            text: "·"
                            font.pointSize: 9
                            opacity: 0.4
                        }
                        ImageButton {
                            source: liked ? "images/liked.png" : "images/like.png"
                            animationDuration: 200
                            sourceSize.height: 16
                            opacity: liked ? 1.0 : 0.3
                            onClicked: function () {
                                if (liked) {
                                    uiBridge.unlikeButton(message.messageid)
                                    liked = false
                                } else {
                                    uiBridge.likeButton(message.messageid)
                                    liked = true
                                }
                            }
                        }
                        Label {
                            text: message.likescount
                            font.pointSize: 9
                            opacity: 0.4
                        }
                        ImageButton {
                            source: "images/menu.png"
                            visible: accountBridge.username == message.author
                            animationDuration: 200
                            sourceSize.height: 16
                            opacity: liked ? 1.0 : 0.3
                            onClicked: function() {
                                postMenu.open()
                            }
                            Menu {
                                id: postMenu
                                x: -width + parent.width
                                y: parent.height
                                transformOrigin: Menu.TopRight

                                MenuItem {
                                    text: qsTr("Delete")
                                    onTriggered: function() {
                                        ComponentCreator.createDeletePopup(this, message).open();
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }
}
