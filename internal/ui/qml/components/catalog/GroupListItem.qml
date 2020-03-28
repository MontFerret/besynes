import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13
import QtQuick.Layouts 1.12
import QtGraphicalEffects 1.12
import "../common" as Common

Control {
    property var model: ({ id: "", name: "", description: "" })
    signal selected(string id)
    signal edited(string id)
    signal deleted(string id)

    id: root

    Button {
        anchors.fill: parent
        flat: true
        onClicked: {
            if (root.selected) {
                root.selected(root.model.id)
            }
        }
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

            Rectangle {
                Layout.fillWidth: true
                Layout.fillHeight: true
                Layout.alignment: Qt.AlignLeft | Qt.AlignVCenter
                color: "transparent"

                ColumnLayout {
                    anchors.fill: parent
                    Text {
                        color: Material.color(Material.Grey, Material.Shade900)
                        font.pixelSize: 14
                        font.family: "Roboto"
                        font.weight: Font.Bold
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
            }

            Common.Dropdown {
                Layout.alignment: Qt.AlignRight
                model: [
                    "Edit",
                    "Delete"
                ]
                onSelected: (option) => {
                    const opt = option.toLowerCase();
                    if (opt === "edit") {
                        if (root.edited) {
                            root.edited(root.model.id)
                        }
                    } else if (opt === "Delete") {
                        if (root.deleted) {
                            root.deleted(root.model.id)
                        }
                    }
                }
            }
        }
    }
}
