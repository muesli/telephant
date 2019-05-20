import QtQuick 2.5
import QtQuick.Controls 2.1
import QtGraphicalEffects 1.0

Image {
    id: img
    property var onClicked: function () {}
    property int roundness: 0
    property int animationDuration: 500

    fillMode: Image.Pad
    horizontalAlignment: Image.AlignHCenter
    verticalAlignment: Image.AlignVCenter
    opacity: 0.3
    smooth: true

    states: State {
        name: "mouse-over"
        when: mouseArea.containsMouse
        PropertyChanges {
            target: img
            opacity: 1.0
        }
    }

    transitions: Transition {
        NumberAnimation {
            properties: "opacity"
            easing.type: Easing.InOutQuad
            duration: animationDuration
        }
    }

    MouseArea {
        id: mouseArea
        anchors.fill: parent
        hoverEnabled: true
        cursorShape: Qt.PointingHandCursor

        onClicked: parent.onClicked()
    }

    layer.enabled: roundness > 0
    layer.effect: OpacityMask {
        maskSource: Item {
            width: img.width
            height: img.height
            Rectangle {
                anchors.fill: parent
                radius: roundness
            }
        }
    }
}
