import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.13

Item {
    property string text: ""
    property bool readOnly: false
    property string placeholder: ""
    property string color: Material.color(Material.Grey, Material.Shade900)
    signal editingFinished(string text)
    readonly property var copy: () => {
        editor.selectAll()
        editor.copy()
        editor.deselect()
    }

    id: root
    anchors.fill: parent

    ScrollView {
        ScrollBar.horizontal.policy: ScrollBar.AlwaysOff
        ScrollBar.vertical.policy: ScrollBar.AsNeeded

        anchors.fill: parent
        clip: true

        TextArea {
            Layout.fillHeight: true
            Layout.fillWidth: true
            id: editor
            text: root.text
            enabled: root.enabled
            readOnly: root.readOnly
            color: root.color
            placeholderText: root.placeholder
            selectionColor: Material.color(Material.Purple)
            mouseSelectionMode: TextEdit.SelectCharacters
            focus: true
            selectByMouse: true
            wrapMode: TextEdit.WordWrap
            padding: 5
            font.pixelSize: 14
            background: Rectangle {
                width: parent.width
                height: parent.height
                border.width: 1
                border.color: Material.color(Material.Grey, Material.Shade200)
                radius: 5
                color: root.enabled ? "white" : Material.color(Material.Grey, Material.Shade100)
            }
            onEditingFinished: root.editingFinished(editor.text)
        }
    }
}
