import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13
import QtQuick.Layouts 1.12

Control {
    property string label: ""
    property string value: ""
    property string placeholder: ""
    signal textChanged(string text)

    id: root

    Grid {
        width: parent.width
        columns: 1

        Label {
            width: root.width
            font.bold: true
            text: root.label
            color: Material.color(Material.Grey, Material.Shade900)
        }

        TextField {
            id: input
            width: root.width
            text: root.value
            selectByMouse: true
            placeholderText: root.placeholder
            color: Material.color(Material.Grey, Material.Shade900)
            selectionColor: Material.color(Material.Purple)
            Material.accent: Material.color(Material.Purple)
            Material.primary: Material.color(Material.Grey, Material.Shade200)
            onTextChanged: (text) => {
                if (root.textChanged) {
                    root.textChanged(input.getText(0, input.length))
                }
            }
        }
    }
}
