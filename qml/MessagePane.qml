import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

ColumnLayout {
    property int idx
    property string name
    property variant messageModel

    MessageList {
        Layout.fillHeight: true
        Layout.fillWidth: true

        id: messagePane
        anchors.margins: 16
        model: messageModel

        headerPositioning: ListView.OverlayHeader

        header: Item {
            z: 2
            width: parent.width
            height: 36

            Label {
                z: 3
                anchors.fill: parent
                anchors.leftMargin: 8
                text: name
                font.pointSize: 15
                font.weight: Font.Light
                verticalAlignment: Label.AlignVCenter
            }
            Label {
                z: 3
                anchors.fill: parent
                anchors.rightMargin: 24
                text: "Close"
                font.pointSize: 10
                font.weight: Font.Light
                horizontalAlignment: Label.AlignRight
                verticalAlignment: Label.AlignVCenter

                MouseArea {
                    anchors.fill: parent
                    cursorShape: Qt.PointingHandCursor
                    onClicked: {
                        uiBridge.closePane(idx)
                    }
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
