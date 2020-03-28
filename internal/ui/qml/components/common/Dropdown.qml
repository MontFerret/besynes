import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.13

RoundButton {
    id: root
    icon.source: "../../icons/more_vert.svg"
    flat: true
    onClicked: menu.open()

    property var model: []

    Menu {
        id: menu
        font.family: "Roboto"
        font.pixelSize: 14

        Repeater {
            model: root.model

            MenuItem {
                text: modelData
            }
        }
    }
}
