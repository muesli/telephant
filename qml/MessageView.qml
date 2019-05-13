import QtQuick 2.5
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3
import QtGraphicalEffects 1.0

ColumnLayout {
    property bool fadeMedia
    property bool showActionButtons: true
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
                    accountPopup.open()
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
                    accountPopup.open()
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
            rounded: true
            opacity: 1.0

            onClicked: function() {
                uiBridge.loadAccount(message.authorid)
                accountPopup.open()
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
                            accountPopup.open()
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
                            accountPopup.open()
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
                            accountPopup.open()
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
                            accountPopup.open()
                        }
                    }
                }
                Label {
                    font.pointSize: 9
                    opacity: 0.4
                    text: message.createdat
                    Layout.fillWidth: true
                    horizontalAlignment: Text.AlignRight

                    MouseArea {
                        anchors.fill: parent
                        cursorShape: Qt.PointingHandCursor
                        onClicked: {
                            Qt.openUrlExternally(message.posturl)
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
                    visible: message.body.length > 0
                    text: "<style>a:link { visibility: hidden; text-decoration: none; color: " + Material.accent + "; }</style>" + message.body
                    textFormat: Text.RichText
                    font.pointSize: 11
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                    opacity: (accountBridge.username == message.author && (message.like || message.forward)) ? 0.4 : 1.0
                    onLinkActivated: function(link) {
                        if (link.startsWith("telephant://")) {
                            var us = link.substr(12, link.length).split("/")

                            if (us[1] == "user") {
                                uiBridge.loadAccount(us[us.length-1])
                                accountPopup.open()
                            }
                            if (us[1] == "tag") {
                                uiBridge.tag(us[us.length-1])
                            }
                        } else
                            Qt.openUrlExternally(link)
                    }

                    MouseArea {
                        anchors.fill: parent
                        // we don't want to eat clicks on the Label
                        acceptedButtons: parent.hoveredLink ? Qt.NoButton : Qt.LeftButton
                        cursorShape: Qt.PointingHandCursor

                        onClicked: function() {
                            uiBridge.loadConversation(message.messageid)
                            conversationPopup.open()
                        }
                    }
                }

                ImageButton {
                    visible: message.mediapreview != ""
                    Layout.topMargin: 4
                    Layout.fillWidth: true
                    // Layout.maximumWidth: sourceSize.width
                    Layout.maximumHeight: (accountBridge.username == message.author && (message.like || message.forward)) ?
                        Math.min(384 / 3, paintedHeight + 8) :
                        Math.min(384, paintedHeight + 8)
                    source: message.mediapreview
                    fillMode: Image.PreserveAspectFit
                    verticalAlignment: Image.AlignBottom
                    autoTransform: true
                    opacity: fadeMedia ? 0.2 : 1.0

                    onClicked: function() {
                        Qt.openUrlExternally(message.mediaurl)
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
                            sourceSize.height: 20
                            onClicked: function () {
                                messagePopup.message = model
                                messagePopup.open()
                            }
                        }
                        ImageButton {
                            source: "images/share.png"
                            animationDuration: 200
                            sourceSize.height: 20
                            opacity: shared ? 1.0 : 0.3
                            onClicked: function () {
                                if (shared) {
                                    uiBridge.unshareButton(message.messageid)
                                    shared = false
                                } else {
                                    sharePopup.message = message
                                    sharePopup.open()
                                }
                            }
                        }
                        ImageButton {
                            source: liked ? "images/liked.png" : "images/like.png"
                            animationDuration: 200
                            sourceSize.height: 20
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
                    }
                }
            }
        }
    }
}
