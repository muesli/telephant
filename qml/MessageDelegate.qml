import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3
import QtGraphicalEffects 1.0

ColumnLayout {
    property bool fadeMedia

    id: messageDelegate
    x: messagePane.Material.elevation
    width: parent.width - messagePane.Material.elevation * 2 - 12
    Pane {
        anchors.horizontalCenter: parent.horizontalCenter
        id: messagePane
        Material.elevation: 6
        Layout.fillWidth: true

        MessageView {
            id: messageView
            fadeMedia: messageDelegate.fadeMedia
            anchors.left: parent.left
            anchors.right: parent.right
        }
    }
}
