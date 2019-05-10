import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
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
            visible: false

            Label {
                text: qsTr("Theme:")
            }

            ComboBox {
                id: themeBox
                property int themeIndex: -1
                Layout.fillWidth: true
                model: ["System", "Material", "Imagine", "Universal", "Light"]
                Component.onCompleted: {
                    themeIndex = find(settings.theme, Qt.MatchFixedString)
                    if (themeIndex !== -1)
                        currentIndex = themeIndex
                }
            }
        }
        RowLayout {
            spacing: 10
            visible: themeBox.currentIndex == 1

            Label {
                text: qsTr("Style:")
            }

            ComboBox {
                id: styleBox
                property int styleIndex: -1
                Layout.fillWidth: true
                model: ["Light", "Dark"]
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
            visible: themeBox.currentIndex !== themeBox.themeIndex ? 1.0 : 0.0
        }

        RowLayout {
            spacing: 10

            Button {
                id: okButton
                Layout.preferredWidth: 0
                Layout.fillWidth: true
                highlighted: true

                text: qsTr("Ok")
                onClicked: {
                    settings.theme = themeBox.displayText
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
                    themeBox.currentIndex = themeBox.themeIndex
                    styleBox.currentIndex = styleBox.styleIndex
                    settingsDialog.close()
                }
            }
        }
    }
}
