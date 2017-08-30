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

        header: Item {
            SystemPalette { id: headerPalette; colorGroup: SystemPalette.Active }
            z: 2
            width: parent.width
            height: 36
            Label {
                z: 3
                anchors.fill: parent
                anchors.leftMargin: 10
                text: name
                font.pixelSize: 18
                font.weight: Font.Light
                verticalAlignment: Label.AlignVCenter
                color: headerPalette.windowText
            }

            Rectangle {
                anchors.fill: parent
                color: headerPalette.window
                opacity: 0.8
            }
        }
    }
}
