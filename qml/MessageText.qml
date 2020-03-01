import QtQuick 2.5
import QtQuick.Controls 2.2
import QtQuick.Controls.Material 2.1

import "componentCreator.js" as ComponentCreator

TextEdit
{
    property var onClicked: function () {}

    id: label
    selectByMouse: true
    readOnly: true
    persistentSelection: true
    selectionColor: Material.accent
    textFormat: Text.RichText

    onLinkActivated: function(link) {
        if (link.startsWith("telephant://")) {
            var us = link.substr(12, link.length).split("/")
            if (us[1] == "user") {
                uiBridge.loadAccount(us[us.length-1])
                ComponentCreator.createAccountPopup(mainWindow).open();
            }
            if (us[1] == "tag") {
                uiBridge.tag(us[us.length-1])
            }
        } else
            Qt.openUrlExternally(link)
    }

    MouseArea {
        anchors.fill: parent
        // we don't want to eat clicks on the Label
        acceptedButtons: Qt.RightButton
        cursorShape: label.hoveredLink ? Qt.PointingHandCursor : Qt.IBeamCursor
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
