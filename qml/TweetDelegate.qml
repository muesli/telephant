import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3
import QtGraphicalEffects 1.0

RowLayout {
    width: parent.width - messagePane.Material.elevation * 2
    Pane {
        id: messagePane
        Material.elevation: 6
        anchors.fill: parent

        MessageView {
            id: messageView
            anchors.left: parent.left
            anchors.right: parent.right
        }
    }
}
