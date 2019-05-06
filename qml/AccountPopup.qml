import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

Popup {
    id: accountPopup
    property string accountid

    modal: true
    // focus: true
    closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutsideParent

    Rectangle {
        anchors.fill: parent
        color: "transparent"

        AccountSummary {
            profile: profileBridge
            id: summary
            anchors.left: parent.left
            anchors.right: parent.right
        }
        ToolSeparator {
            id: separator
            anchors.top: summary.bottom
            anchors.left: parent.left
            anchors.right: parent.right
            orientation: Qt.Horizontal
        }
        MessageList {
            anchors.top: separator.bottom
            anchors.bottom: parent.bottom
            anchors.left: parent.left
            anchors.right: parent.right

            fadeMedia: false
            model: accountBridge.accountMessages
        }
    }
}
