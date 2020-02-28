import QtQuick 2.5
import QtQuick.Controls 2.2
import QtQuick.Controls.Material 2.1

TextEdit
{
    property var onClicked: function () {}

    id: label
    selectByMouse: true
    readOnly: true
    persistentSelection: true
    selectionColor: Material.accent

    MouseArea {
        id: ma1
        anchors.fill: parent
        // we don't want to eat clicks on the Label
        acceptedButtons: Qt.RightButton
        cursorShape: parent.hoveredLink ? Qt.PointingHandCursor : Qt.IBeamCursor
        hoverEnabled: true
        propagateComposedEvents: true

        onReleased: {
            if (mouse.button == Qt.RightButton) {
                contextMenu.x = mouse.x;
                contextMenu.y = mouse.y;
                contextMenu.open();
                return;
            }

            mouse.accepted = false;
        }

        Menu {
            id: contextMenu
            MenuItem {
                text: "Copy"
                onTriggered: label.copy();
            }
        }
    }
}
