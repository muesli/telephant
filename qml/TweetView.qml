import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

ColumnLayout {
    property string name
    property variant messageModel

    TweetList {
        Layout.fillHeight: true
        Layout.fillWidth: true

        id: tweetView
        anchors.margins: 16
        model: messageModel

        headerPositioning: ListView.OverlayHeader

        header: Rectangle {
            width: parent.width
            height: 32
            z: 2
            color: Material.background

            Label {
                anchors.fill: parent
                text: name
                font.pixelSize: 18
                font.weight: Font.Light
            }
        }
    }
}
