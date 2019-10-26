import QtQuick 2.13
import "../common" as Common

Item {
    property string values: ""

    id: root

    Common.CodeEditor {
        anchors.fill: parent
        text: root.values
        placeholder: "Parameter values"
    }
}
