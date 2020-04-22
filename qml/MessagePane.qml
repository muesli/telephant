import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.11

ColumnLayout {
    property int idx
    property string name
    property bool sticky
    property variant messageModel

    MessageList {
        Layout.fillHeight: true
        Layout.fillWidth: true
        Layout.minimumWidth: 360

        id: messagePane
        anchors.margins: 16
        model: messageModel

        headerPositioning: ListView.OverlayHeader

        header: Item {
            z: 2
            width: parent.width
            height: 36

            Label {
                anchors.left: parent.left
                anchors.top: parent.top
                anchors.bottom: parent.bottom
                anchors.leftMargin: 8
                z: 3
                text: name
                font.pointSize: 15
                font.weight: Font.Light
                verticalAlignment: Label.AlignVCenter
            }
            TextButton {
                anchors.right: parent.right
                anchors.top: parent.top
                anchors.bottom: parent.bottom
                anchors.leftMargin: 8
                anchors.rightMargin: 24
                anchors.bottomMargin: 4
                visible: !sticky
                z: 3
                text: "Close"
                font.pointSize: 10
                font.weight: Font.Light
                verticalAlignment: Label.AlignVCenter

                onClicked: function() {
                    uiBridge.closePane(idx)
                }
            }

            Pane {
                anchors.fill: parent
                opacity: 0.8

                background: Rectangle {
                    color: Material.color(Material.Grey, Material.Shade900)
                }
            }
        }
    }
}
