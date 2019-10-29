import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13
import QtQuick.Layouts 1.12
import "../common" as Common

Common.Dialog {
    id: root
    title: "Settings"
    anchors.centerIn: parent
    width: 400
    height: 200
    padding: 15

    GridLayout {
        anchors.fill: parent
        columns: 1

        GeneralForm {
            Layout.fillHeight: true
            Layout.fillWidth: true
        }
    }

    RowLayout {
        width: parent.width
        anchors.bottom: parent.bottom

        Button {
            Layout.alignment: Qt.AlignLeft
            Material.background: Material.color(Material.Grey, Material.Shade300)
            Material.foreground: Material.color(Material.Grey, Material.Shade900)
            text: "Cancel"
            onClicked: {
                root.reject()
            }
        }

        Button {
            Layout.alignment: Qt.AlignRight
            Material.background: Material.Indigo
            Material.foreground: Material.color(Material.Grey, Material.Shade50)
            text: "Save"
            onClicked: {
                root.accept()
            }
        }
    }
}
