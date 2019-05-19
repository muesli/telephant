import QtQuick 2.4

ListModel {
    Component.onCompleted: {
        append({"panename": "Messages", "msgmodel": messageModel});
        append({"panename": "Notifications", "panesticky": true, "msgmodel": notificationModel});
    }
}
