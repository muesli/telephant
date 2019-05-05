import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

Popup {
    id: conversationPopup
    property string messageid

    modal: true
    // focus: true
    closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutsideParent

    ColumnLayout {
        anchors.fill: parent

        MessagePane {
            Layout.row: 0
            Layout.column: 0

            name: qsTr("Conversation")
            messageModel: accountBridge.conversation
        }
    }
}
