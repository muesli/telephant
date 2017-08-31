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
    minimumHeight: 480

    Item {
        TweetPopup {
            id: tweetPopup
            width: 480
            height: 300
            x: mainWindow.width / 2 - width / 2
            y: mainWindow.height / 2 - height / 2 - mainWindow.header.height / 2
        }

        AboutDialog {
            id: aboutDialog
            x: (mainWindow.width - width) / 2
            y: mainWindow.height / 6
            width: Math.min(mainWindow.width, mainWindow.height) / 3 * 2
        }

        SettingsDialog {
            id: settingsDialog
            x: (mainWindow.width - width) / 2
            y: mainWindow.height / 6
            width: Math.min(mainWindow.width, mainWindow.height) / 3 * 2
        }
    }

    Component {
        id: messagesDelegate

        TweetDelegate { }
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
                        Qt.openUrlExternally(
                                    "https://twitter.com/" + accountBridge.username)
                    }
                }

                Label {
                    id: titleLabel
                    text: accountBridge.username
                    font.pixelSize: 16
                    elide: Label.ElideRight
                    horizontalAlignment: Image.AlignHCenter
                    verticalAlignment: Qt.AlignVCenter
                }
            }

            ToolButton {
                id: menuButton
                anchors.right: parent.right
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

        ListView {
            id: listView
            currentIndex: -1
            anchors.fill: parent

            delegate: ItemDelegate {
                width: parent.width
                text: model.title
                highlighted: ListView.isCurrentItem
                onClicked: {
                    listView.currentIndex = -1
                    drawer.close()

                    switch (model.sid) {
                    case 0:
                        tweetPopup.messageid = 0
                        tweetPopup.open()
                        break
                    case 1:
                        Qt.quit()
                        break
                    }
                }
            }

            model: ListModel {
                ListElement {
                    title: qsTr("New Tweet")
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

    GridLayout {
        id: maingrid
        columns: 2
        rows: 1
        anchors.fill: parent
        anchors.margins: 0
        columnSpacing: 0
        rowSpacing: 0

        TweetView {
            Layout.row: 0
            Layout.column: 0

            name: qsTr("Messages")
            messageModel: accountBridge.messages
        }

        TweetView {
            Layout.row: 0
            Layout.column: 1

            name: qsTr("Notifications")
            messageModel: accountBridge.notifications
        }
    }
}
