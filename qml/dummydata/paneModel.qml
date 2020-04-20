import QtQuick 2.13

ListModel {
    Component.onCompleted: {
        append({"panename": "Messages", "msgmodel": messageModel});
        append({"panename": "Notifications", "panesticky": true, "msgmodel": notificationModel});
    }
}
