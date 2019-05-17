import QtQuick 2.4
import QtQuick.Controls 2.5
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

ApplicationWindow {
    id: mainWindow
    visible: true

    Material.theme: settings.style == "Dark" ? Material.Dark : Material.Light
    Material.accent: Material.Purple
    // flags: Qt.FramelessWindowHint
    background: Rectangle {
        color: Material.color(Material.Grey, Material.Shade900)
    }

    minimumWidth: 364
    minimumHeight: 590
    width: minimumWidth * 2
    height: minimumWidth * 1.5

    Component.onCompleted: {
        if (settings.firstRun) {
            connectDialog.open()
        }
    }

    Item {
        AboutDialog {
            id: aboutDialog
            x: (mainWindow.width - width) / 2
            y: mainWindow.height / 6
            width: Math.min(mainWindow.width, mainWindow.height) / 3 * 2
        }

        ConnectDialog {
            id: connectDialog
            x: mainWindow.width / 2 - width / 2
            y: mainWindow.height / 2 - height / 2 - mainWindow.header.height
            width: 340
            height: 500
        }

        SettingsDialog {
            id: settingsDialog
            x: (mainWindow.width - width) / 2
            y: mainWindow.height / 6
            width: Math.min(mainWindow.width, mainWindow.height) / 3 * 2
        }

        Popup {
            id: errorDialog
            modal: true
            focus: true
            contentHeight: errorLayout.height
            visible: accountBridge.error.length > 0
            x: mainWindow.width / 2 - width / 2
            y: mainWindow.height / 2 - height / 2 - mainWindow.header.height
            width: Math.min(mainWindow.width * 0.66, errorLayout.implicitWidth + 32)

            ColumnLayout {
                id: errorLayout
                spacing: 20
                width: parent.width

                Label {
                    text: qsTr("Error")
                    font.bold: true
                }

                Label {
                    wrapMode: Label.Wrap
                    font.pointSize: 14
                    text: accountBridge.error
                }

                Button {
                    id: okButton
                    Layout.alignment: Qt.AlignCenter
                    highlighted: true

                    text: qsTr("Close")
                    onClicked: {
                        accountBridge.error = ""
                        errorDialog.close()
                    }
                }
            }
        }
    }

    header: ToolBar {
        ToolButton {
            id: drawerButton
            contentItem: Image {
                fillMode: Image.Pad
                horizontalAlignment: Image.AlignHCenter
                verticalAlignment: Image.AlignVCenter
                source: "images/drawer.png"
            }
            onClicked: {
                drawer.open()
            }
        }

        RowLayout {
            anchors.verticalCenter: parent.verticalCenter
            anchors.horizontalCenter: parent.horizontalCenter
            spacing: 8

            ImageButton {
                opacity: 1.0
                rounded: true
                horizontalAlignment: Image.AlignHCenter
                verticalAlignment: Image.AlignVCenter
                source: accountBridge.avatar
                sourceSize.height: 32
                onClicked: function() {
                    Qt.openUrlExternally(accountBridge.profileURL)
                }
            }

            Label {
                id: titleLabel
                text: accountBridge.username
                font.pointSize: 13
                elide: Label.ElideRight
                horizontalAlignment: Image.AlignHCenter
                verticalAlignment: Qt.AlignVCenter
            }
        }

        ToolButton {
            id: postButton
            anchors.right: menuButton.left
            contentItem: Image {
                fillMode: Image.Pad
                horizontalAlignment: Image.AlignHCenter
                verticalAlignment: Image.AlignVCenter
                source: "images/post.png"
            }
            onClicked: {
                messagePopup.message = null
                messagePopup.open()
            }
        }
        ToolButton {
            anchors.right: parent.right
            id: menuButton
            Layout.alignment: Qt.AlignRight
            contentItem: Image {
                fillMode: Image.Pad
                horizontalAlignment: Image.AlignHCenter
                verticalAlignment: Image.AlignVCenter
                source: "images/menu.png"
            }
            onClicked: optionsMenu.open()

            Menu {
                id: optionsMenu
                x: parent.width - width
                transformOrigin: Menu.TopRight

                MenuItem {
                    text: qsTr("Connect")
                    onTriggered: function() {
                        connectDialog.reset()
                        connectDialog.open()
                    }
                }
                /*
                MenuItem {
                    text: qsTr("Settings")
                    onTriggered: settingsDialog.open()
                }
                */
                MenuItem {
                    text: qsTr("About")
                    onTriggered: aboutDialog.open()
                }
            }
        }
    }

    Drawer {
        id: drawer
        width: drawerLayout.implicitWidth + 16
        height: mainWindow.height
        dragMargin: 0

        ColumnLayout {
            id: drawerLayout
            anchors.fill: parent

            AccountSummary {
                profile: accountBridge
            }
            ToolSeparator {
                Layout.fillWidth: true
                orientation: Qt.Horizontal
            }

            ListView {
                id: listView
                currentIndex: -1
                Layout.fillWidth: true
                Layout.fillHeight: true
                delegate: ItemDelegate {
                    width: parent.width
                    text: model.title
                    highlighted: ListView.isCurrentItem
                    onClicked: {
                        listView.currentIndex = -1
                        drawer.close()
                        switch (model.sid) {
                        case 0:
                            messagePopup.message = null
                            messagePopup.open()
                            break
                        case 1:
                            Qt.quit()
                            break
                        }
                    }
                }
                model: ListModel {
                    ListElement {
                        title: qsTr("New Post")
                        property int sid: 0
                    }
                    ListElement {
                        title: qsTr("Exit")
                        property int sid: 1
                    }
                }
                ScrollIndicator.vertical: ScrollIndicator {
                }
            }
        }
    }

    ScrollView {
        id: mainscroll
        anchors.fill: parent
        ScrollBar.horizontal.policy: contentWidth > width ? ScrollBar.AlwaysOn : ScrollBar.AlwaysOff
        ScrollBar.vertical.policy: ScrollBar.AlwaysOff
        contentWidth: Math.max(maingrid.implicitWidth, parent.width)

        GridLayout {
            id: maingrid
            // columns: accountBridge.panes.length
            rows: 1
            anchors.fill: parent
            anchors.margins: 0
            columnSpacing: 0
            rowSpacing: 0

            Repeater {
                model: accountBridge.panes
                MessagePane {
                    Layout.row: 0
                    Layout.column: index

                    idx: index
                    name: model.panename
                    sticky: model.panesticky
                    messageModel: model.msgmodel
                }
            }
        }
    }
}
