import QtQuick 2.4
import QtQuick.Controls 2.2
import QtQuick.Controls.Material 2.2
import QtQuick.Layouts 1.3

Popup {
    id: popup
    property string name
    property string messageid
    property string author
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

    modal: true
    // focus: true
    closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutsideParent

    ColumnLayout {
        anchors.fill: parent

        MessageView {
            visible: popup.messageid > 0
            name: popup.name
            messageid: popup.messageid
            author: popup.author
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
        }

        Label {
            visible: popup.messageid > 0
            text: qsTr("Replying to %1").arg(name)
            opacity: 0.3
        }

        TextArea {
            id: messageArea
            Layout.fillWidth: true
            Layout.fillHeight: true
            focus: true
            placeholderText: popup.messageid > 0 ? qsTr("Tweet your reply") : qsTr(
                                                       "What's happening?")
            wrapMode: TextArea.Wrap
        }

        RowLayout {
            anchors.fill: parent

            Label {
                id: remCharsLabel

                anchors.verticalCenter: sendButton.verticalCenter
                anchors.right: sendButton.left
                anchors.rightMargin: 12

                font.pixelSize: 16
                text: 140 - messageArea.text.length
            }

            Button {
                id: sendButton
                enabled: remCharsLabel.text >= 0 && messageArea.text.length > 0
                anchors.bottom: parent.bottom
                anchors.right: parent.right
                highlighted: true
                Material.accent: Material.Blue
                text: popup.messageid > 0 ? qsTr("Reply") : qsTr("Tweet")

                onClicked: {
                    popup.close()
                    uiBridge.tweetButton(popup.messageid,
                                         "@" + name + " " + messageArea.text)
                    messageArea.clear()
                }
            }
        }
    }
}
