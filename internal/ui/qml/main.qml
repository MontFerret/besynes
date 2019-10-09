import QtQuick 2.13
import QtQuick.Controls 2.13
import "./components/query"

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
