import QtQuick 2.12
import QtQuick.Controls 2.12

ListView {
    property bool fadeMedia: true

    id: view
    spacing: 12
    clip: true
    cacheBuffer: 10000

    ScrollBar.vertical: ScrollBar {
        width: 12
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

    ScrollHelper {
        id: scrollHelper
        flickable: view
        anchors.fill: view
    }
}
