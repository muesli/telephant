import QtQuick 2.4

QtObject {
    property string username: "username"
    property string name: "Name"
    property string avatar: "https://pbs.twimg.com/profile_images/908139250612363264/m-CkMJbl_400x400.jpg"
    property int posts: 1337
    property int follows: 123
    property int followers: 42
    property int postSizeLimit: 500

    property ListModel messages: messageModel
}
