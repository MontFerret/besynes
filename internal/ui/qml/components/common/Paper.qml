import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13

Pane {
    Material.elevation: 6
    background: Rectangle {
        anchors.fill: parent
        color: "white"
        border.color: Material.color(Material.Grey, Material.Shade300)
    }
}
