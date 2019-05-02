import QtQuick 2.5
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3
import QtGraphicalEffects 1.0

RowLayout {
    property string name: model.name
    property string messageid: model.messageid
    property string posturl: model.posturl
    property string author: model.author
    property string authorurl: model.authorurl
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
    property string media: model.media

    RowLayout {
        spacing: 8
        ImageButton {
            id: image
            anchors.top: parent.top
            source: avatar
            sourceSize.width: 48
            width: 48
            fillMode: Image.PreserveAspectCrop
            roundness: 8
            rounded: true
            opacity: 0.8
        }
        ColumnLayout {
            Layout.fillWidth: true
            spacing: 4

            RowLayout {
                anchors.left: parent.left
                anchors.right: parent.right

                Label {
                    id: namelabel
                    font.bold: true
                    text: name
                    textFormat: Text.PlainText
                    Layout.fillWidth: true
                    Layout.maximumWidth: implicitWidth + 1
                    elide: Text.ElideRight
                }
                Label {
                    // anchors.bottom: parent.bottom
                    color: "steelblue"
                    font.pixelSize: 11
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
                    font.pixelSize: 11
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
                width: parent.width
                // anchors.bottom: parent.bottom
                spacing: 4
                Label {
                    visible: body.length > 0
                    text: body
                    textFormat: Text.RichText
                    onLinkActivated: Qt.openUrlExternally(link)
                    font.pixelSize: 13
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap

                    MouseArea {
                        anchors.fill: parent
                        acceptedButtons: Qt.NoButton // we don't want to eat clicks on the Label
                        cursorShape: parent.hoveredLink ? Qt.PointingHandCursor : Qt.ArrowCursor
                    }
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
                                cursorShape: Qt.PointingHandCursor
                                onClicked: {
                                    Qt.openUrlExternally(posturl)
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
                            text: qsTr("%1 shared").arg(actorname)
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

                        ImageButton {
                            source: "images/reply.svg"
                            sourceSize.height: 16
                            onClicked: function () {
                                messagePopup.name = name
                                messagePopup.messageid = messageid
                                messagePopup.author = author
                                messagePopup.avatar = avatar
                                messagePopup.body = body
                                messagePopup.createdat = createdat
                                messagePopup.posturl = posturl
                                messagePopup.actor = actor
                                messagePopup.actorname = actorname
                                messagePopup.reply = reply
                                messagePopup.replytoid = replytoid
                                messagePopup.replytoauthor = replytoauthor
                                messagePopup.forward = forward
                                messagePopup.mention = mention
                                messagePopup.like = like
                                messagePopup.open()
                            }
                        }
                        ImageButton {
                            source: "images/retweet.svg"
                            sourceSize.height: 16
                            onClicked: function () {
                                uiBridge.shareButton(messageid)
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

                ImageButton {
                    visible: media != ""
                    Layout.fillWidth: true
                    // Layout.maximumWidth: sourceSize.width
                    Layout.maximumHeight: Math.min(384, paintedHeight + 8)
                    source: media
                    fillMode: Image.PreserveAspectFit
                    verticalAlignment: Image.AlignBottom
                    autoTransform: true
                    opacity: 0.2

                    onClicked: function() {
                        Qt.openUrlExternally(media)
                    }
                }
            }
        }
    }
}
