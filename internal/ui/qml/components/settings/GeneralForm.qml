import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13
import QtQuick.Layouts 1.12
import "../common" as Common
import "../common/forms" as Forms

Control {
    width: parent.width
    height: parent.height

    Grid {
        anchors.fill: parent

        Forms.TextInput {
            width: parent.width
            label: "CDP"
            value: "http://127.0.0.1:9222"
            placeholder: "CDP Address"
        }
    }
}
