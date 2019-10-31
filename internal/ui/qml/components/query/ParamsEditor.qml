import QtQuick 2.13
import "../common" as Common

Item {
    property string values: ""
    signal editingFinished(string text)

    id: root

    Common.CodeEditor {
        anchors.fill: parent
        text: root.values
        placeholder: "Parameter values"
        enabled: root.enabled
        readOnly: !root.enabled
        onEditingFinished: function (text) {
            root.editingFinished(text)
        }
    }
}
