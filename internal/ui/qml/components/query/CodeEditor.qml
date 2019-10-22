import QtQuick 2.13

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
            anchors.fill: parent
            padding: 10
            color: "black"
            focus: true
            enabled: root.enabled
            selectionColor: "green"
            mouseSelectionMode: TextEdit.SelectWords
            text: text
            wrapMode: TextEdit.WrapAnywhere
            onEditingFinished: {
                root.text = textEditor.text
            }
        }
    }
}
