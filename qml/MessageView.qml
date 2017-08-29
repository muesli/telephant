import QtQuick 2.4
import QtQuick.Controls 2.2
import QtQuick.Controls.Material 2.2
import QtQuick.Layouts 1.3
import QtGraphicalEffects 1.0

RowLayout {
    property string name: model.name
    property string messageid: model.messageid
    property string author: model.author
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

    RowLayout {
        spacing: 8
        ImageButton {
            id: image
            anchors.top: parent.top
            source: avatar
            sourceSize.width: 48
            fillMode: Image.PreserveAspectCrop
            roundness: 8
            rounded: true
            opacity: 0.8
        }
        ColumnLayout {
            anchors.top: parent.top
            Layout.fillWidth: true
            spacing: 4

            RowLayout {
                anchors.top: parent.top
                anchors.left: parent.left
                anchors.right: parent.right

                Label {
                    id: namelabel
                    font.bold: true
                    text: name
                }
                Label {
                    // anchors.bottom: parent.bottom
                    color: "steelblue"
                    font.pixelSize: 11
                    text: "@" + author

                    MouseArea {
                        anchors.fill: parent
                        onClicked: {
                            Qt.openUrlExternally(
                                        "https://twitter.com/" + author)
                        }
                    }
                }
                Label {
                    anchors.right: parent.right
                    font.pixelSize: 11
                    text: createdat

                    MouseArea {
                        anchors.fill: parent
                        onClicked: {
                            Qt.openUrlExternally(
                                        "https://twitter.com/statuses/" + messageid)
                        }
                    }
                }
            }
            ColumnLayout {
                width: parent.width
                anchors.bottom: parent.bottom
                spacing: 4
                Label {
                    text: body
                    textFormat: Text.RichText
                    onLinkActivated: Qt.openUrlExternally(link)
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                }
                RowLayout {
                    anchors.left: parent.left
                    anchors.right: parent.right
                    RowLayout {
                        visible: reply
                        Label {
                            font.pixelSize: 12
                            text: qsTr("Replying to %1").arg(
                                      "@" + replytoauthor)
                            opacity: 0.3

                            MouseArea {
                                anchors.fill: parent
                                onClicked: {
                                    Qt.openUrlExternally(
                                                "https://twitter.com/statuses/" + replytoid)
                                }
                            }
                        }
                    }
                    RowLayout {
                        visible: forward && !like
                        Image {
                            smooth: true
                            source: "images/retweet.svg"
                            sourceSize.height: 14
                            opacity: 0.5
                        }
                        Label {
                            font.pixelSize: 12
                            text: qsTr("%1 retweeted").arg(actorname)
                            opacity: 0.3
                        }
                    }
                    RowLayout {
                        visible: like
                        Image {
                            smooth: true
                            source: "images/like.svg"
                            sourceSize.height: 14
                            opacity: 0.5
                        }
                        Label {
                            font.pixelSize: 12
                            text: qsTr("%1 liked").arg(actorname)
                            opacity: 0.3
                        }
                    }

                    RowLayout {
                        anchors.right: parent.right
                        /* Button {
                                    highlighted: true
                                    Material.accent: Material.Green
                                    text: qsTr("Reply")
                                    onClicked: {
                                        tweetPopup.open()
                                        // tweetModel.setData(tweetModel.index(index, 0) , true, "editing");
                                    }
                                } */
                        ImageButton {
                            source: "images/reply.svg"
                            sourceSize.height: 16
                            onClicked: function () {
                                tweetPopup.name = name
                                tweetPopup.messageid = messageid
                                tweetPopup.author = author
                                tweetPopup.avatar = avatar
                                tweetPopup.body = body
                                tweetPopup.createdat = createdat
                                tweetPopup.actor = actor
                                tweetPopup.actorname = actorname
                                tweetPopup.reply = reply
                                tweetPopup.replytoid = replytoid
                                tweetPopup.replytoauthor = replytoauthor
                                tweetPopup.forward = forward
                                tweetPopup.mention = mention
                                tweetPopup.like = like
                                tweetPopup.open()
                            }
                        }
                        ImageButton {
                            source: "images/retweet.svg"
                            sourceSize.height: 16
                            onClicked: function () {
                                uiBridge.retweetButton(messageid)
                            }
                        }
                        ImageButton {
                            source: "images/like.svg"
                            sourceSize.height: 16
                            onClicked: function () {
                                uiBridge.likeButton(messageid)
                            }
                        }
                    }
                }
            }
        }
    }
}
