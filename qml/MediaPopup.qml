import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3

Popup {
    property var url

    id: popup

    modal: true
    focus: true
    height: image.height + 16
    width: image.width + 16
    anchors.centerIn: mainWindow.overlay
    closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutsideParent

    Image {
        id: image
        height: Math.min(sourceSize.height, mainWindow.height * 0.8)
        width: Math.min(sourceSize.width, mainWindow.width * 0.8)
        anchors.centerIn: parent
        smooth: true
        fillMode: Image.PreserveAspectFit
        source: url
    }
}
