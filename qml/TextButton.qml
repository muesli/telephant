import QtQuick 2.5
import QtQuick.Controls 2.1

Label {
    id: label
    property var onClicked: function () {}

    MouseArea {
        id: mouseArea
        anchors.fill: parent
        hoverEnabled: true
        cursorShape: Qt.PointingHandCursor

        onClicked: parent.onClicked()
    }
}
