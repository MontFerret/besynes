import QtQuick 2.13

Item {
    id: root

    property var value: ({ data: "", error: "" })

    Rectangle {
        anchors.fill: parent
        border.width: 1
        border.color: "#EEEEEE"
        radius: 5

        Text {
            id: results
            anchors.fill: parent
            color: value.error ? "red" : "black"
            focus: true
            text: value.error ? value.error : value.data
            padding: 10
            wrapMode: TextEdit.WrapAnywhere
        }
    }
}
