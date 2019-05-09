import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

Popup {
    id: popup
    property string name
    property string messageid
    property string posturl
    property string author
    property string authorurl
    property string authorid
    property string avatar
    property string body
    property string createdat
    property string actor
    property string actorname
    property bool reply
    property string replytoid
    property string replytoauthor
    property bool forward
    property bool mention
    property bool like
    property string mediapreview
    property string mediaurl
    property bool liked
    property bool shared

    modal: true
    // focus: true
    height: Math.min(mainWindow.height * 0.8, layout.implicitHeight + 32)
    closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutsideParent

    Flickable {
        id: flickable
        anchors.fill: parent
        clip: true
        contentHeight: layout.height

        ColumnLayout {
            id: layout
            width: parent.width
            MessageView {
                visible: popup.messageid > 0
                name: popup.name
                messageid: popup.messageid
                posturl: popup.posturl
                author: popup.author
                authorid: popup.authorid
                authorurl: popup.authorurl
                avatar: popup.avatar
                body: popup.body
                createdat: popup.createdat
                actor: popup.actor
                actorname: popup.actorname
                reply: popup.reply
                replytoid: popup.replytoid
                replytoauthor: popup.replytoauthor
                forward: popup.forward
                mention: popup.mention
                like: popup.like
                mediapreview: popup.mediapreview
                mediaurl: popup.mediaurl
                liked: popup.liked
                shared: popup.shared
            }

            Label {
                visible: popup.messageid > 0
                text: qsTr("Replying to %1").arg(name)
                opacity: 0.3
            }

            TextArea {
                id: messageArea
                Layout.fillWidth: true
                Layout.minimumHeight: 128
                focus: true
                selectByMouse: true
                placeholderText: popup.messageid > 0 ? qsTr("Post your reply") : qsTr("What's happening?")
                wrapMode: TextArea.Wrap
            }

            RowLayout {
                Layout.alignment: Qt.AlignRight

                Label {
                    id: remCharsLabel

                    Layout.alignment: Qt.AlignVCenter | Qt.AlignRight

                    font.pixelSize: 16
                    text: accountBridge.postSizeLimit - messageArea.text.length
                }

                Button {
                    id: sendButton
                    enabled: remCharsLabel.text >= 0 && messageArea.text.length > 0
                    Layout.alignment: Qt.AlignBottom | Qt.AlignRight
                    highlighted: true
                    // Material.accent: Material.Blue
                    text: popup.messageid > 0 ? qsTr("Reply") : qsTr("Post")

                    onClicked: {
                        popup.close()
                        var msg = messageArea.text
                        if (popup.messageid > 0) {
                            msg = "@" + author + " " + msg
                        }
                        uiBridge.postButton(popup.messageid, msg)
                        messageArea.clear()
                    }
                }
            }
        }
    }
}
