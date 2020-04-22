import QtQuick 2.12
import QtQuick.Controls 2.12

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
