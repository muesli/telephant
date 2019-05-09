import QtQuick 2.5
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3
import QtGraphicalEffects 1.0

ColumnLayout {
    property bool fadeMedia
    property bool showActionButtons: true

    property string name: model.name
    property string messageid: model.messageid
    property string posturl: model.posturl
    property string author: model.author
    property string authorurl: model.authorurl
    property string authorid: model.authorid
    property string avatar: model.avatar
    property string body: model.body
    property string createdat: model.createdat
    property string actor: model.actor
    property string actorname: model.actorname
    property bool reply: model.reply
    property string replytoid: model.replytoid
    property string replytoauthor: model.replytoauthor
    property bool forward: model.forward
    property bool mention: model.mention
    property bool like: model.like
    property bool followed: model.followed
    property bool following: model.following
    property bool followedby: model.followedby
    property string mediapreview: model.mediapreview
    property string mediaurl: model.mediaurl
    property bool liked: model.liked
    property bool shared: model.shared

    RowLayout {
        visible: forward && !like
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
            text: qsTr("%1 shared").arg(actorname)
            opacity: (accountBridge.username == author && (like || forward)) ? 0.8 : 0.3
        }
    }
    RowLayout {
        visible: like
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
            text: qsTr("%1 liked").arg(actorname)
            opacity: (accountBridge.username == author && (like || forward)) ? 0.8 : 0.3
        }
    }

    RowLayout {
        spacing: 8

        ImageButton {
            id: image
            Layout.alignment: Qt.AlignTop
            source: avatar
            sourceSize.width: 48
            width: 48
            fillMode: Image.PreserveAspectCrop
            roundness: 4
            rounded: true
            opacity: 1.0

            onClicked: function() {
                uiBridge.loadAccount(authorid)
                accountPopup.open()
            }
        }
        RowLayout {
            visible: followed
            Layout.fillWidth: true
            spacing: 4

            ColumnLayout {
                Layout.fillWidth: true
                Label {
                    font.pointSize: 11
                    font.bold: true
                    text: qsTr("%1 followed you").arg(actorname)
                    textFormat: Text.PlainText
                    Layout.fillWidth: true
                    elide: Text.ElideRight
                    opacity: 1.0
                }
                Label {
                    font.pointSize: 11
                    text: actor
                    textFormat: Text.PlainText
                    Layout.fillWidth: true
                    elide: Text.ElideRight
                    opacity: 1.0
                }
            }
            Button {
                Layout.alignment: Qt.AlignVCenter | Qt.AlignRight
                highlighted: true
                text: following ? qsTr("Unfollow") : qsTr("Follow")

                onClicked: {
                    uiBridge.followButton(authorid, !following)
                }
            }
        }
        ColumnLayout {
            visible: !followed
            Layout.fillWidth: true
            spacing: 4

            RowLayout {
                anchors.left: parent.left
                anchors.right: parent.right

                Label {
                    id: namelabel
                    font.pointSize: 11
                    font.bold: true
                    text: name
                    textFormat: Text.PlainText
                    Layout.fillWidth: true
                    Layout.maximumWidth: implicitWidth + 1
                    elide: Text.ElideRight
                    opacity: (accountBridge.username == author && (like || forward)) ? 0.4 : 1.0
                }
                Label {
                    // anchors.bottom: parent.bottom
                    font.pointSize: 9
                    opacity: 0.4
                    text: "@" + author

                    MouseArea {
                        anchors.fill: parent
                        cursorShape: Qt.PointingHandCursor
                        onClicked: {
                            Qt.openUrlExternally(authorurl)
                        }
                    }
                }
                Label {
                    anchors.right: parent.right
                    font.pointSize: 9
                    opacity: 0.4
                    text: createdat

                    MouseArea {
                        anchors.fill: parent
                        cursorShape: Qt.PointingHandCursor
                        onClicked: {
                            Qt.openUrlExternally(posturl)
                        }
                    }
                }
            }
            ColumnLayout {
                visible: !followed
                // width: parent.width
                // anchors.bottom: parent.bottom
                spacing: 4
                Label {
                    visible: body.length > 0
                    text: "<style>a:link { visibility: hidden; text-decoration: none; color: " + Material.accent + "; }</style>" + body
                    textFormat: Text.RichText
                    onLinkActivated: Qt.openUrlExternally(link)
                    font.pointSize: 11
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                    opacity: (accountBridge.username == author && (like || forward)) ? 0.4 : 1.0

                    MouseArea {
                        anchors.fill: parent
                        // we don't want to eat clicks on the Label
                        acceptedButtons: parent.hoveredLink ? Qt.NoButton : Qt.LeftButton
                        cursorShape: Qt.PointingHandCursor

                        onClicked: function() {
                            uiBridge.loadConversation(messageid)
                            conversationPopup.open()
                        }
                    }
                }

                ImageButton {
                    visible: mediapreview != ""
                    Layout.topMargin: 4
                    Layout.fillWidth: true
                    // Layout.maximumWidth: sourceSize.width
                    Layout.maximumHeight: (accountBridge.username == author && (like || forward)) ?
                        Math.min(384 / 3, paintedHeight + 8) :
                        Math.min(384, paintedHeight + 8)
                    source: mediapreview
                    fillMode: Image.PreserveAspectFit
                    verticalAlignment: Image.AlignBottom
                    autoTransform: true
                    opacity: fadeMedia ? 0.2 : 1.0

                    onClicked: function() {
                        Qt.openUrlExternally(mediaurl)
                    }
                }

                RowLayout {
                    anchors.left: parent.left
                    anchors.right: parent.right
                    RowLayout {
                        visible: reply
                        Label {
                            font.pointSize: 10
                            text: qsTr("Replying to %1").arg(
                                      "@" + replytoauthor)
                            opacity: 0.4

                            MouseArea {
                                anchors.fill: parent
                                cursorShape: Qt.PointingHandCursor
                                onClicked: {
                                    Qt.openUrlExternally(posturl)
                                }
                            }
                        }
                    }

                    RowLayout {
                        visible: showActionButtons && !(accountBridge.username == author && (like || forward))
                        anchors.right: parent.right
                        Layout.topMargin: 4

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
                                    uiBridge.unshareButton(messageid)
                                    shared = false
                                } else {
                                    sharePopup.message = model
                                    sharePopup.open()
                                }
                            }
                        }
                        ImageButton {
                            source: liked ? "images/liked.png" : "images/like.png"
                            animationDuration: 200
                            sourceSize.height: 20
                            onClicked: function () {
                                if (liked) {
                                    uiBridge.unlikeButton(messageid)
                                    liked = false
                                } else {
                                    uiBridge.likeButton(messageid)
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
