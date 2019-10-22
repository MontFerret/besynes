import QtQuick 2.13

Item {
    property string value: qsTr("")

    Rectangle {
        anchors.fill: parent
        border.width: 1
        border.color: "#EEEEEE"
        radius: 5

        Text {
            id: results
            anchors.fill: parent
            color: "black"
            focus: true
            text: value
            padding: 10
            wrapMode: TextEdit.WrapAnywhere
        }
    }
}
