import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13
import QtQuick.Layouts 1.12
import QtGraphicalEffects 1.12
import "../common" as Common

Control {
    property var model: ({ id: "", name: "", description: "", queries: [] })

    id: root
    state: "folded"
    states: [
        State {
            name: "folded"

        },
        State {
            name: "unfolded"
        }
    ]

    Button {
        anchors.fill: parent
        flat: true
        contentItem: RowLayout {
            Rectangle {
                Layout.alignment: Qt.AlignLeft
                Layout.minimumHeight: 50
                Layout.minimumWidth: 50
                Layout.preferredHeight: 50
                Layout.preferredWidth: 50
                color: "transparent"

                Image {
                    anchors.centerIn: parent
                    id: iconFolder
                    source: "../../icons/folder.svg"
                    width: 24
                    height: 24
                }

                ColorOverlay {
                    anchors.fill: iconFolder
                    source: iconFolder
                    color: Material.color(Material.Grey, Material.Shade700)
                }
            }
            ColumnLayout {
                Layout.fillWidth: true
                Layout.fillHeight: true
                Text {
                    color: Material.color(Material.Grey, Material.Shade800)
                    font.pixelSize: 14
                    font.family: "Roboto"
                    font.bold: true
                    antialiasing: true
                    text: root.model.name
                }

                Text {
                    color: Material.color(Material.Grey, Material.Shade700)
                    font.pixelSize: 12
                    font.family: "Roboto"
                    antialiasing: true
                    text: root.model.queries.count + " queries"
                }
            }

            RoundButton {
                Layout.alignment: Qt.AlignRight
                icon.source: "../../icons/more_vert.svg"
                flat: true
            }
        }
    }
}