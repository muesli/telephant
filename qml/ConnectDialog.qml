import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

Popup {
    id: connectDialog
    property string instance

    modal: true
    focus: true
    closePolicy: Popup.CloseOnEscape

    property var reset: function() {
        connectSwipeView.currentIndex = 0
        instanceArea.text = ""
        codeArea.text = ""
    }

    ColumnLayout {
        spacing: 16
        anchors.fill: parent
        clip: true

        Label {
            text: qsTr("Add an Account")
            Layout.alignment: Qt.AlignHCenter
            font.bold: true
        }

        Image {
            id: logo
            Layout.alignment: Qt.AlignHCenter
            smooth: true
            source: "images/accounts/mastodon.svg"
            sourceSize.height: 96
        }

        SwipeView {
            id: connectSwipeView
            Layout.fillWidth: true
            Layout.fillHeight: true
            Layout.alignment: Qt.AlignHCenter
            Component.onCompleted: contentItem.interactive = false

            currentIndex: 0
            Item {
                id: instancePage

                ColumnLayout {
                        anchors.fill: parent
                        anchors.margins: 16

                        TextField {
                            id: instanceArea
                            focus: true
                            selectByMouse: true
                            placeholderText: qsTr("Instance, e.g. https://mastodon.social")
                            Layout.fillWidth: true
                        }

                        Button {
                            id: connectButton
                            enabled: instanceArea.text.length > 0
                            Layout.alignment: Qt.AlignBottom | Qt.AlignCenter
                            highlighted: true
                            text: qsTr("Authorize Telephant")

                            onClicked: {
                                var instance = instanceArea.text
                                var result = uiBridge.connectButton(instance)
                                if (result) {
                                    connectSwipeView.currentIndex = 1
                                }
                            }
                        }
                }
            }

            Item {
                id: authPage

                ColumnLayout {
                        anchors.fill: parent
                        anchors.margins: 16

                        Label {
                            text: "You need to retrieve an authorization code from your instance:"
                            Layout.alignment: Qt.AlignCenter
                            Layout.fillWidth: true
                            wrapMode: Text.WordWrap
                        }
                        TextArea {
                            id: authURL
                            Layout.fillWidth: true
                            text: settings.authURL
                            readOnly: true
                        }
                        RowLayout {
                            Layout.alignment: Qt.AlignHCenter
                            Button {
                                highlighted: true
                                text: qsTr("Open in Browser")
                                onClicked: {
                                    Qt.openUrlExternally(settings.authURL)
                                }
                            }
                            Button {
                                highlighted: true
                                text: qsTr("Copy URL")
                                onClicked: {
                                    authURL.selectAll()
                                    authURL.copy()
                                }
                            }
                        }
                        Item {
                            height: 16
                        }

                        TextField {
                            id: codeArea
                            focus: true
                            selectByMouse: true
                            placeholderText: qsTr("Auth code provided by your instance")
                            Layout.fillWidth: true
                        }

                        Button {
                            id: authButton
                            enabled: codeArea.text.length > 0
                            Layout.alignment: Qt.AlignBottom | Qt.AlignCenter
                            highlighted: true
                            text: qsTr("Login")

                            onClicked: {
                                var code = codeArea.text
                                var result = uiBridge.authButton(code, settings.redirectURL)
                                if (result) {
                                    connectDialog.close()
                                }
                            }
                        }
                }
            }
        }

        PageIndicator {
            id: indicator
            Layout.alignment: Qt.AlignHCenter

            count: connectSwipeView.count
            currentIndex: connectSwipeView.currentIndex
        }
    }
}
