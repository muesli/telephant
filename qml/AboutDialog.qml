import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
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
                text: "Telephant! by <a style=\"text-decoration: none; color: orange;\" href=\"https://twitter.com/mueslix\">@mueslix</a>"
                textFormat: Text.RichText
                wrapMode: Label.Wrap
                font.pixelSize: 12
                onLinkActivated: Qt.openUrlExternally(link)

                MouseArea {
                    anchors.fill: parent
                    acceptedButtons: Qt.NoButton // we don't want to eat clicks on the Label
                    cursorShape: parent.hoveredLink ? Qt.PointingHandCursor : Qt.ArrowCursor
                }
            }

            Label {
                width: aboutDialog.availableWidth
                text: qsTr("Telephant! is a light-weight but modern social media client")
                textFormat: Text.RichText
                wrapMode: Label.Wrap
                font.pixelSize: 12
                onLinkActivated: Qt.openUrlExternally(link)

                MouseArea {
                    anchors.fill: parent
                    acceptedButtons: Qt.NoButton // we don't want to eat clicks on the Label
                    cursorShape: parent.hoveredLink ? Qt.PointingHandCursor : Qt.ArrowCursor
                }
            }
        }
    }
