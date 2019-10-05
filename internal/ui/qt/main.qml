import QtQuick 2.13
import QtQuick.Controls 2.5
import QtQuick.Controls.Styles 1.4
import QtQuick.Controls.Material 2.12

ApplicationWindow {
    id: win
    visible: true
    width: 1024
    height: 768
    title: qsTr("Besynes")

    background: Rectangle {
        color: "#EEEEEE"
        anchors.fill: parent
    }

    QueryTabView {
        id: tabs
    }
}
