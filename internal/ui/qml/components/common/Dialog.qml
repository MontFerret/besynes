import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13
import QtQuick.Layouts 1.12

Dialog {
    id: root
    modal: true

    header: ToolBar {
        Material.background: Material.DeepPurple

        Label {
            text: root.title
            anchors.centerIn: parent
            color: "white"
        }
    }


    background: Rectangle {
        color: Material.color(Material.Grey, Material.Shade200)
    }

    footer: Rectangle {
        color: Material.color(Material.Grey, Material.Shade200)
    }
}
