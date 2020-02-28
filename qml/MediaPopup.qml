import QtQuick 2.4
import QtQuick.Controls 2.1
import QtQuick.Controls.Material 2.1
import QtQuick.Layouts 1.3
import QtMultimedia 5.9

Popup {
    property var url

    id: popup

    modal: true
    focus: true
    height: mediaItem.height + 16
    width: mediaItem.width + 16
    anchors.centerIn: mainWindow.overlay

    Item {
        id: mediaItem
        anchors.centerIn: parent
        width: image.visible ?
            image.width :
            video.width

        height: image.visible ?
            image.height :
            video.height

        Image {
            id: image
            width: Math.min(sourceSize.width, mainWindow.width * 0.8)
            height: Math.min(sourceSize.height, mainWindow.height * 0.8)
            anchors.centerIn: parent
            smooth: true
            fillMode: Image.PreserveAspectFit
            source: visible ? url : ""
            visible: !(url.endsWith(".webm") || url.endsWith(".mp4"))
        }

        Video {
            id: video
            width: metaData.resolution ? Math.min(metaData.resolution.width, mainWindow.width * 0.8) : 0
            height: metaData.resolution ? Math.min(metaData.resolution.height, mainWindow.height * 0.8) : 0
            autoLoad: true
            autoPlay: false
            loops: MediaPlayer.Infinite
            anchors.centerIn: parent
            fillMode: VideoOutput.PreserveAspectFit
            source: visible ? url : ""
            visible: url.endsWith(".webm") || url.endsWith(".mp4")

            onStatusChanged: {
                if (status == MediaPlayer.Loaded)
                    video.play()
            }

/*
            MouseArea {
                anchors.fill: parent
                onClicked: {
                    if (video.playbackState != MediaPlayer.PausedState) {
                        video.play()
                    } else {
                        video.pause()
                    }
                }
            }
*/
        }
    }
}
