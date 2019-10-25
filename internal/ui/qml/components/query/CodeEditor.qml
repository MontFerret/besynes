import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.13

Item {
    property string text: qsTr("")
    property bool enabled: true
    property string background: "white"

    id: root

    Rectangle {
        anchors.fill: parent
        border.width: 1
        border.color: "#EEEEEE"
        radius: 5
        color: background

        TextEdit {
            id: textEditor
            text: text
            anchors.fill: parent
            padding: 10
            color: Material.color(Material.Grey, Material.Shade900)
            focus: true
            enabled: root.enabled
            selectByMouse: true
            mouseSelectionMode: TextEdit.SelectCharacters
            selectionColor: Material.color(Material.Purple)
            wrapMode: TextEdit.WordWrap
            onEditingFinished: {
                root.text = textEditor.text
            }
        }
    }
}
