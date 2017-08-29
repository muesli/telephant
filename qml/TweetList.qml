import QtQuick 2.4
import QtQuick.Controls 2.1

ListView {
    id: view
    spacing: 12
    clip: true

    ScrollBar.vertical: ScrollBar {
        width: 8
    }

    delegate: messagesDelegate

    Label {
        anchors.fill: parent
        horizontalAlignment: Qt.AlignHCenter
        verticalAlignment: Qt.AlignVCenter
        visible: parent.count == 0
        text: qsTr("No messages to show yet!")
        font.bold: true
    }
}
