import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.11
import QtGraphicalEffects 1.12

ColumnLayout {
    property bool fadeMedia

    id: messageDelegate
    x: messagePane.Material.elevation
    width: parent.width - messagePane.Material.elevation * 2 - 12
    Pane {
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
