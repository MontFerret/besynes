import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Styles 1.4
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.13

Rectangle {
    function isVertical() {
        return parent.orientation === Qt.Vertical
    }

    id: root
    implicitWidth: 8
    implicitHeight: 8
    color: Material.color(Material.Grey, Material.Shade50)
    state: SplitHandle.pressed ? "pressed" : "released"
    states: [
        State {
            name: "released"
            PropertyChanges {
                target: handle;
                width: isVertical() ? 10 : 2;
                height: isVertical() ? 2 : 10;
                radius: 3
            }
        },

        State {
            name: "pressed"
            PropertyChanges {
                target: handle;
                width: 3;
                height: 3;
                radius: 3
            }
        }
    ]

    Rectangle {
        id: handle
        anchors.centerIn: parent
        color: Material.color(Material.Grey, Material.Shade700)

        Behavior on height {
            PropertyAnimation {
                easing.type: Easing.InQuad;
                duration: 100
            }
        }

        Behavior on width {
            PropertyAnimation {
                easing.type: Easing.InQuad;
                duration: 100
            }
        }
    }
}
