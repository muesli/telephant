import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

Popup {
    id: connectDialog
    property string instance

    modal: true
    focus: true
    closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutsideParent

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
            sourceSize.height: 128
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
                                connectSwipeView.currentIndex = 1
                                uiBridge.connectButton(instance)
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
                            text: "<a href=\"" + settings.authURL + "\">Click here to request a new authorization code from your instance</a>"
                            Layout.alignment: Qt.AlignCenter
                            Layout.fillWidth: true
                            textFormat: Text.RichText
                            onLinkActivated: Qt.openUrlExternally(link)
                            font.pointSize: 12
                            wrapMode: Text.WordWrap

                            MouseArea {
                                anchors.fill: parent
                                acceptedButtons: Qt.NoButton // we don't want to eat clicks on the Label
                                cursorShape: parent.hoveredLink ? Qt.PointingHandCursor : Qt.ArrowCursor
                            }
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
                                connectDialog.close()
                                uiBridge.authButton(code, settings.redirectURL)
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

            // anchors.bottom: connectSwipeView.bottom
            // anchors.horizontalCenter: parent.horizontalCenter
        }

    }
}
