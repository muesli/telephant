import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

Popup {
    property var message

    id: popup

    modal: true
    focus: true
    height: Math.min(mainWindow.height * 0.8, layout.implicitHeight + 32)
    width: Math.min(mainWindow.width * 0.66, 500)
    anchors.centerIn: mainWindow.overlay
    closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutsideParent

    Flickable {
        id: flickable
        anchors.fill: parent
        clip: true
        contentHeight: layout.height

        BusyIndicator {
            z: 1
            id: busy
            running: false
            anchors.centerIn: parent
        }

        DropArea {
            id: drop
            anchors.fill: parent
            enabled: true

            onEntered:
                console.log("entered")

            onExited:
                console.log("exited")

            onDropped: {
                console.log("dropped", drop.urls.length, "urls")

                for (var i = 0; i < drop.urls.length; i++) {
                    console.log(drop.urls[i])

                    busy.running = true
                    var media = uiBridge.uploadAttachment(drop.urls[i])
                    /*if (media != '') {
                        attachments.append({"id": media, "url": drop.urls[i]})
                    }*/
                }
                drop.acceptProposedAction()
            }
        }

        ColumnLayout {
            id: layout
            width: parent.width

            MessageView {
                showActionButtons: false
                visible: message != null
                message: popup.message
            }

            Label {
                visible: message != null
                text: qsTr("Replying to %1").arg(message.name)
                opacity: 0.3
            }

            TextArea {
                id: messageArea
                Layout.fillWidth: true
                Layout.minimumHeight: 128
                focus: true
                selectByMouse: true
                placeholderText: message != null ? qsTr("Post your reply") : qsTr("What's happening?")
                wrapMode: TextArea.Wrap
            }

            Connections {
                target: accountBridge.attachments
                onRowsInserted: {
                    busy.running = false
                }
                onRowsRemoved: {
                }
            }

            Flow {
                id: attachmentLayout
                Layout.fillWidth: true
                Repeater {
                    model: accountBridge.attachments
                    Image {
                        smooth: true
                        source: model.attachmentPreview
                        sourceSize.height: 64

                        MouseArea {
                            anchors.fill: parent
                            cursorShape: Qt.PointingHandCursor

                            onClicked: function() {
                                accountBridge.attachments.removeAttachment(index)
                            }
                        }
                    }
                }
            }

            RowLayout {
                Layout.alignment: Qt.AlignRight

                Label {
                    id: remCharsLabel

                    Layout.alignment: Qt.AlignVCenter | Qt.AlignRight

                    font.pointSize: 12
                    text: accountBridge.postSizeLimit - uiBridge.postLimitCount(messageArea.text)
                }

                Button {
                    id: sendButton
                    enabled: remCharsLabel.text >= 0 && messageArea.text.length > 0
                    Layout.alignment: Qt.AlignBottom | Qt.AlignRight
                    highlighted: true
                    text: message != null ? qsTr("Reply") : qsTr("Post")

                    onClicked: {
                        popup.close()
                        var msg = messageArea.text
                        var msgid = "";
                        if (message != null) {
                            msgid = message.messageid
                            msg = "@" + message.author + " " + msg
                        }
                        uiBridge.postButton(msgid, msg)
                        messageArea.clear()
                    }
                }
            }
        }
    }
}
