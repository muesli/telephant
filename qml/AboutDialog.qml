import QtQuick 2.4
import QtQuick.Controls 2.2
import QtQuick.Controls.Material 2.2
import QtQuick.Layouts 1.3

    Popup {
        modal: true
        focus: true

        contentHeight: aboutColumn.height

        Column {
            id: aboutColumn
            spacing: 20

            Label {
                text: qsTr("About")
                font.bold: true
            }

            Label {
                width: aboutDialog.availableWidth
                text: "Chirp! by <a style=\"text-decoration: none; color: orange;\" href=\"https://twitter.com/mueslix\">@mueslix</a>"
                textFormat: Text.RichText
                wrapMode: Label.Wrap
                font.pixelSize: 12
                onLinkActivated: Qt.openUrlExternally(link)
            }

            Label {
                width: aboutDialog.availableWidth
                text: qsTr("Chirp! is a light-weight but modern Twitter client")
                textFormat: Text.RichText
                wrapMode: Label.Wrap
                font.pixelSize: 12
                onLinkActivated: Qt.openUrlExternally(link)
            }
        }
    }
