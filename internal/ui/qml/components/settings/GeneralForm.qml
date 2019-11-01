import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13
import QtQuick.Layouts 1.12
import "../common" as Common
import "../common/forms" as Forms

Item {
    property string cdpAddress: ""
    signal changed(string text)

    id: root
    width: parent.width
    height: parent.height

    Grid {
        anchors.fill: parent

        Forms.TextInput {
            enabled: root.enabled
            width: parent.width
            label: "CDP"
            value: root.cdpAddress
            placeholder: "CDP Address"
            onTextChanged: (text) => {
                if (root.changed) {
                    root.changed(text)
                }
            }
        }
    }
}
