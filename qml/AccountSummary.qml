import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

ColumnLayout {
    RowLayout {
        id: accountLayout
        anchors.horizontalCenter: parent.horizontalCenter
        spacing: 16
        Layout.topMargin: 16
        Layout.leftMargin: 16
        Layout.bottomMargin: 4
        ImageButton {
            opacity: 1.0
            rounded: true
            horizontalAlignment: Image.AlignHCenter
            verticalAlignment: Image.AlignVCenter
            source: accountBridge.avatar
            sourceSize.height: 64
            onClicked: function() {
                Qt.openUrlExternally(accountBridge.profileURL)
            }
        }
        ColumnLayout {
            Layout.fillWidth: true
            Label {
                Layout.fillWidth: true
                Layout.alignment: Qt.AlignLeft
                text: accountBridge.name
                font.pixelSize: 16
                font.bold: true
                elide: Label.ElideRight
            }
            Label {
                text: accountBridge.username
                font.pixelSize: 16
                opacity: 0.7
                elide: Label.ElideRight
            }
        }
    }
    RowLayout {
        anchors.horizontalCenter: parent.horizontalCenter
        Label {
            Layout.alignment: Qt.AlignLeft
            text: "<b>" + accountBridge.posts + "</b> Posts"
            font.pixelSize: 11
            elide: Label.ElideRight
        }
        Label {
            Layout.alignment: Qt.AlignCenter
            text: "<b>" + accountBridge.follows + "</b> Follows"
            font.pixelSize: 11
            elide: Label.ElideRight
        }
        Label {
            Layout.alignment: Qt.AlignRight
            text: "<b>" + accountBridge.followers + "</b> Followers"
            font.pixelSize: 11
            elide: Label.ElideRight
        }
    }
}
