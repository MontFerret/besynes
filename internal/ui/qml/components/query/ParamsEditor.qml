import QtQuick 2.13

Item {
    property string values: qsTr("")

    Rectangle {
        anchors.fill: parent
        border.width: 1
        border.color: "#EEEEEE"
        radius: 5

        TextEdit {
            id: jsonEditor
            anchors.fill: parent
            color: "black"
            focus: true
            text: text
            padding: 10
            wrapMode: TextEdit.WrapAnywhere
        }
    }
}
