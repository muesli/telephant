import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

ApplicationWindow {
    id: mainWindow
    visible: true
    Material.theme: Material.Dark
    Material.accent: Material.Purple

    // flags: Qt.FramelessWindowHint
    minimumWidth: 800
    minimumHeight: 450

    Item {
        MessagePopup {
            id: messagePopup
            width: mainWindow.width * 0.66
            // height: mainWindow.height * 0.8
            x: mainWindow.width / 2 - width / 2
            y: mainWindow.height / 2 - height / 2 - mainWindow.header.height
        }

        SharePopup {
            id: sharePopup
            width: mainWindow.width * 0.66
            // height: mainWindow.height * 0.8
            x: mainWindow.width / 2 - width / 2
            y: mainWindow.height / 2 - height / 2 - mainWindow.header.height
        }

        ConversationPopup {
            id: conversationPopup
            width: mainWindow.width * 0.8
            height: mainWindow.height * 0.8
            x: mainWindow.width / 2 - width / 2
            y: mainWindow.height / 2 - height / 2 - mainWindow.header.height
        }

        AccountPopup {
            id: accountPopup
            width: mainWindow.width * 0.8
            height: mainWindow.height * 0.8
            x: mainWindow.width / 2 - width / 2
            y: mainWindow.height / 2 - height / 2 - mainWindow.header.height
        }

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
            height: 340
        }

        SettingsDialog {
            id: settingsDialog
            x: (mainWindow.width - width) / 2
            y: mainWindow.height / 6
            width: Math.min(mainWindow.width, mainWindow.height) / 3 * 2
        }
    }

    header: ToolBar {
        /* MouseArea {
            anchors.fill: parent;
            property variant clickPos: "1,1"

            onPressed: {
                clickPos  = Qt.point(mouse.x,mouse.y)
            }

            onPositionChanged: {
                var delta = Qt.point(mouse.x-clickPos.x, mouse.y-clickPos.y)
                mainWindow.x += delta.x;
                mainWindow.y += delta.y;
            }
        } */

        RowLayout {
            spacing: 20
            anchors.fill: parent

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
                Layout.topMargin: 4
                Layout.bottomMargin: 4

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
                        text: qsTr("Settings")
                        onTriggered: settingsDialog.open()
                    }
                    MenuItem {
                        text: qsTr("About")
                        onTriggered: aboutDialog.open()
                    }
                }
            }
        }
    }

    Drawer {
        id: drawer
        width: mainWindow.width / 3
        height: mainWindow.height
        dragMargin: 0

        ColumnLayout {
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
                            connectDialog.open()
                            break
                        case 1:
                            messagePopup.message = null
                            messagePopup.open()
                            break
                        case 2:
                            Qt.quit()
                            break
                        }
                    }
                }
                model: ListModel {
                    ListElement {
                        title: qsTr("Connect")
                        property int sid: 0
                    }
                    ListElement {
                        title: qsTr("New Post")
                        property int sid: 1
                    }
                    ListElement {
                        title: qsTr("Exit")
                        property int sid: 2
                    }
                }
                ScrollIndicator.vertical: ScrollIndicator {
                }
            }
        }
    }

    GridLayout {
        id: maingrid
        columns: 2
        rows: 1
        anchors.fill: parent
        anchors.margins: 0
        columnSpacing: 0
        rowSpacing: 0

        MessagePane {
            Layout.row: 0
            Layout.column: 0

            name: qsTr("Messages")
            messageModel: accountBridge.messages
        }

        MessagePane {
            Layout.row: 0
            Layout.column: 1

            name: qsTr("Notifications")
            messageModel: accountBridge.notifications
        }
    }
}
