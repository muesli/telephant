import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.11

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

            RowLayout {
                Image {
                    Layout.leftMargin: 8
                    smooth: true
                    source: "images/telephant_logo.png"
                    sourceSize.height: 64
                }

                ColumnLayout {
                    Layout.leftMargin: 8
                    Label {
                        text: "<a style=\"text-decoration: none; color: white;\" href=\"https://mastodon.social/@telephant\">Telephant</a>"
                        textFormat: Text.RichText
                        wrapMode: Label.Wrap
                        font.pointSize: 14
                        onLinkActivated: Qt.openUrlExternally(link)

                        MouseArea {
                            anchors.fill: parent
                            acceptedButtons: Qt.NoButton // we don't want to eat clicks on the Label
                            cursorShape: parent.hoveredLink ? Qt.PointingHandCursor : Qt.ArrowCursor
                        }
                    }

                    Label {
                        text: "Version 0.1"
                        textFormat: Text.RichText
                        wrapMode: Label.Wrap
                        font.pointSize: 10
                        onLinkActivated: Qt.openUrlExternally(link)

                        MouseArea {
                            anchors.fill: parent
                            acceptedButtons: Qt.NoButton // we don't want to eat clicks on the Label
                            cursorShape: parent.hoveredLink ? Qt.PointingHandCursor : Qt.ArrowCursor
                        }
                    }
                }
            }

            Label {
                width: aboutDialog.availableWidth
                text: qsTr("Telephant is a light-weight but slick Mastodon client")
                textFormat: Text.RichText
                wrapMode: Label.Wrap
                font.pointSize: 12
                onLinkActivated: Qt.openUrlExternally(link)

                MouseArea {
                    anchors.fill: parent
                    acceptedButtons: Qt.NoButton // we don't want to eat clicks on the Label
                    cursorShape: parent.hoveredLink ? Qt.PointingHandCursor : Qt.ArrowCursor
                }
            }
        }
    }
