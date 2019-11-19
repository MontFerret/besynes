import QtQuick 2.0
import QtQuick.Controls.Material 2.13

Item {
    id: root

    property int orientation: Qt.Horizontal
    property int align: Qt.AlignBottom
    property color color: Material.color(Material.Grey, Material.Shade300)
    property real size: parent.width

    Rectangle {
        width: root.orientation === Qt.Horizontal ? root.size : 1
        height: root.orientation === Qt.Horizontal ? 1 : root.size
        anchors.bottom: root.align === Qt.AlignBottom ? parent.bottom : undefined
        anchors.top: root.align === Qt.AlignTop ? parent.top : undefined
        anchors.right: root.align === Qt.AlignRight ? parent.right : undefined
        anchors.left: root.align === Qt.AlignLeft ? parent.left : undefined
        color: root.color
    }
}
