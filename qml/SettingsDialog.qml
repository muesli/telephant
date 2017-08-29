import QtQuick 2.4
import QtQuick.Controls 2.2
import QtQuick.Controls.Material 2.2
import QtQuick.Layouts 1.3

Popup {
    modal: true
    focus: true

    height: settingsColumn.implicitHeight + topPadding + bottomPadding

    contentItem: ColumnLayout {
        id: settingsColumn
        spacing: 20

        Label {
            text: qsTr("Settings")
            font.bold: true
        }

        RowLayout {
            spacing: 10

            Label {
                text: qsTr("Style:")
            }

            ComboBox {
                id: styleBox
                property int styleIndex: -1
                Layout.fillWidth: true
                model: ["Default", "Material", "Universal"]
                Component.onCompleted: {
                    styleIndex = find(settings.style, Qt.MatchFixedString)
                    if (styleIndex !== -1)
                        currentIndex = styleIndex
                }
            }
        }

        Label {
            horizontalAlignment: Label.AlignHCenter
            verticalAlignment: Label.AlignVCenter
            Layout.fillWidth: true
            Layout.fillHeight: true

            text: qsTr("Restart required")
            opacity: styleBox.currentIndex !== styleBox.styleIndex ? 1.0 : 0.0
        }

        RowLayout {
            spacing: 10

            Button {
                id: okButton
                Layout.preferredWidth: 0
                Layout.fillWidth: true

                text: qsTr("Ok")
                onClicked: {
                    settings.style = styleBox.displayText
                    settingsDialog.close()
                }
            }

            Button {
                id: cancelButton
                Layout.preferredWidth: 0
                Layout.fillWidth: true

                text: qsTr("Cancel")
                onClicked: {
                    styleBox.currentIndex = styleBox.styleIndex
                    settingsDialog.close()
                }
            }
        }
    }
}
