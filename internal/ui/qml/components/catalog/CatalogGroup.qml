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
                Layout.minimumHeight: 24
                Layout.minimumWidth: 24
                Layout.preferredHeight: 30
                Layout.preferredWidth: 50
                color: "transparent"

                Image {
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
                    color: Material.color(Material.Grey, Material.Shade700)
                    font.pixelSize: 14
                    font.family: "Roboto"
                    font.bold: true
                    text: root.model.name
                }

                Text {
                    color: Material.color(Material.Grey, Material.Shade700)
                    font.pixelSize: 12
                    font.family: "Roboto"
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
