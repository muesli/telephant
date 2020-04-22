import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.11

Popup {
    id: accountPopup
    property string accountid

    modal: true
    // focus: true
    width: Math.min(mainWindow.width * 0.8, 600)
    height: mainWindow.height * 0.8
    anchors.centerIn: mainWindow.overlay

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
