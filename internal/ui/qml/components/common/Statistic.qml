import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.13

Item {
    id: root

    property string name: ""
    property string value: ""

    Row {
        anchors.fill: parent
        visible: root.visible
        spacing: 4

        Text {
            id: textName
            color: Material.color(Material.Grey, Material.Shade800)
            text: `${root.name}:`
            font.pixelSize: 10
        }

        Text {
            id: textValue
            color: Material.color(Material.Blue)
            text: root.value
            font.pixelSize: 10
        }
    }
}
