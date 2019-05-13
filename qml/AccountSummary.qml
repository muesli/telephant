import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

ColumnLayout {
    property var profile

    RowLayout {
        id: accountLayout
        spacing: 16
        Layout.topMargin: 16
        Layout.leftMargin: 16
        Layout.bottomMargin: 4
        ImageButton {
            height: 64
            width: 64
            opacity: 1.0
            roundness: 4
            rounded: true
            horizontalAlignment: Image.AlignHCenter
            verticalAlignment: Image.AlignVCenter
            source: profile.avatar
            sourceSize.height: 64
            onClicked: function() {
                // Qt.openUrlExternally(profile.profileURL)
                uiBridge.loadAccount(profile.profileID)
                accountPopup.open()
            }
        }
        ColumnLayout {
            Layout.fillWidth: true
            Label {
                Layout.fillWidth: true
                text: profile.name
                font.pointSize: 13
                font.bold: true
                elide: Label.ElideRight
            }
            Label {
                text: profile.username + (profile.followedBy ? " (follows you)" : "")
                font.pointSize: 11
                opacity: 0.7
                elide: Label.ElideRight
            }
        }
        Button {
            id: followButton
            Layout.alignment: Qt.AlignBottom | Qt.AlignRight
            visible: profile.profileID != accountBridge.profileID
            highlighted: true
            text: profile.following ? qsTr("Unfollow") : qsTr("Follow")

            onClicked: {
                uiBridge.followButton(profile.profileID, !profile.following)
            }
        }
    }
    RowLayout {
        Item {
            Layout.fillWidth: true
        }
        Label {
            Layout.alignment: Qt.AlignLeft
            text: "<b>" + profile.posts + "</b> Posts"
            font.pointSize: 10
            elide: Label.ElideRight
        }
        Label {
            Layout.alignment: Qt.AlignCenter
            text: "<b>" + profile.followCount + "</b> Follows"
            font.pointSize: 10
            elide: Label.ElideRight
        }
        Label {
            Layout.alignment: Qt.AlignRight
            text: "<b>" + profile.followerCount + "</b> Followers"
            font.pointSize: 10
            elide: Label.ElideRight
        }
        Item {
            Layout.fillWidth: true
        }
    }
}
