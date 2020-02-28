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

    Flickable {
        id: flickable
        anchors.fill: parent
        clip: true
        contentHeight: layout.height

        ColumnLayout {
            id: layout
            width: parent.width

            MessageView {
                showActionButtons: false
                message: popup.message
            }

            RowLayout {
                Layout.alignment: Qt.AlignRight

                Button {
                    id: sendButton
                    Layout.alignment: Qt.AlignBottom | Qt.AlignRight
                    highlighted: true
                    text: qsTr("Delete")

                    onClicked: {
                        popup.close()
                        uiBridge.deleteButton(message.messageid)
                    }
                }
            }
        }
    }
}
