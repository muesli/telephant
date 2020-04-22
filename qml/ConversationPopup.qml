import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.11

Popup {
    id: conversationPopup
    property string messageid

    modal: true
    focus: true
    width: Math.min(mainWindow.width * 0.8, 600)
    height: mainWindow.height * 0.8
    anchors.centerIn: mainWindow.overlay

    ColumnLayout {
        anchors.fill: parent

        MessageList {
            Layout.fillHeight: true
            Layout.fillWidth: true

            fadeMedia: false
            model: accountBridge.conversation
        }
    }
}
