import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.13
import "../common" as Common

Item {
    property string text: qsTr("")
    signal editingFinished(string text)

    id: root
    anchors.fill: parent

    Common.CodeEditor {
        id: codeEditor
        anchors.fill: parent
        text: root.text
        placeholder: "Query text"
        readOnly: !root.enabled
        enabled: root.enabled
        onEditingFinished: function (text) {
            root.editingFinished(text)
        }
    }
}
