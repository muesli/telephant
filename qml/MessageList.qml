import QtQuick 2.4
import QtQuick.Controls 2.1

ListView {
    property bool fadeMedia: true

    id: view
    spacing: 12
    clip: true

    ScrollBar.vertical: ScrollBar {
        width: 8
        background: Rectangle {
            color: "transparent"
        }
    }

    delegate: MessageDelegate {
        fadeMedia: view.fadeMedia
    }

    Label {
        anchors.fill: parent
        horizontalAlignment: Qt.AlignHCenter
        verticalAlignment: Qt.AlignVCenter
        visible: parent.count == 0
        text: qsTr("No messages to show yet!")
        font.bold: true
    }
}
