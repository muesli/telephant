import QtQuick 2.12

ListModel {
    Component.onCompleted: {
        append({"panename": "Messages", "msgmodel": messageModel});
        append({"panename": "Notifications", "panesticky": true, "msgmodel": notificationModel});
    }
}
