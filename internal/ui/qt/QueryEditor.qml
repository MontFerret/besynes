import QtQuick 2.13
import QtQuick.Controls 2.5
import QtQuick.Layouts 1.13
import QtQuick.Controls.Material 2.12

Item {
    property string text: qsTr("")

    Page {
        anchors.fill: parent
        padding: 15

        header: ToolBar {
            Material.background: Material.DeepPurple

            leftPadding: 15
            rightPadding: 15

            RowLayout {
                anchors.fill: parent
                Label {
                    text: "Title"
                    elide: Label.ElideRight
                    horizontalAlignment: Qt.AlignHCenter
                    verticalAlignment: Qt.AlignVCenter
                    Layout.fillWidth: true
                }
                Button {
                    id: execBtn
                    text: "Exec"
                    highlighted: true
                    Material.background: Material.Blue
                }
            }
        }

        Rectangle {
            anchors.fill: parent
            color: "white"

            TextEdit {
                anchors.fill: parent
                id: code
                color: "black"
                focus: true
                text: text
            }
        }
    }
}
