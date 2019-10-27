import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13
import QtQuick.Layouts 1.12
import "../common" as Common

Common.Dialog {
    id: root
    title: "Settings"
    standardButtons: Dialog.Cancel | Dialog.Save
    anchors.centerIn: parent
    width: 600
    height: 300

    ColumnLayout {
        anchors.fill: parent

        Item {
            Layout.fillWidth: true
            Layout.fillHeight: true

            Label {
                text: "General"
            }

            Common.Paper {
                RowLayout {
                    Label {
                        text: "CDP address"
                    }

                    TextInput {}
                }
            }
        }
    }
}
