import QtQuick 2.13
import QtQuick.Controls 2.13

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
