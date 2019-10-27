import QtQuick 2.0
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13

Rectangle {
    id: root
    color: Material.color(Material.Grey, Material.Shade700)
    anchors.fill: parent
    z: 10
    opacity: 0.8

    BusyIndicator {
        anchors.centerIn: parent
        running: root.visible
        z: 12
    }
}
