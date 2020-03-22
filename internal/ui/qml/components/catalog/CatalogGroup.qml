import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13
import QtQuick.Layouts 1.12
import QtGraphicalEffects 1.12
import "../common" as Common

Control {
    id: root

    property var model: ({ id: "", name: "", description: "", queries: [] })
    state: "folded"
    states: [
        State {
            name: "folded"

        },
        State {
            name: "unfolded"
        }
    ]

    function getBackgroundColor() {
        if (root.state === "folded") {
            return Material.color(Material.Grey, Material.Shade50)
        }

        return Material.color(Material.Grey, Material.Shade300)
    }

    Pane {
        anchors.fill: parent

        background: Rectangle {
            id: background
            color: getBackgroundColor()

            Common.Line {
                width: parent.width
            }

            MouseArea {
                id: mouse
                anchors.fill: parent
                hoverEnabled: true
                onEntered: {
                    if (root.state === "folded") {
                        background.color = Material.color(Material.Grey, Material.Shade300)
                        iconCaretContainer.color = Material.color(Material.Grey, Material.Shade300)
                        iconFolderContainer.color = Material.color(Material.Grey, Material.Shade300)
                        textContainer.color = Material.color(Material.Grey, Material.Shade300)
                    }
                }
                onExited: {
                    if (root.state === "folded") {
                        background.color = Material.color(Material.Grey, Material.Shade50)
                        iconCaretContainer.color = Material.color(Material.Grey, Material.Shade50)
                        iconFolderContainer.color = Material.color(Material.Grey, Material.Shade50)
                        textContainer.color = Material.color(Material.Grey, Material.Shade50)
                    }
                }
                onClicked: {
                    root.state = root.state === "folded" ? "unfolded" : "folded";
                }
            }
        }

        RowLayout {
            anchors.fill: parent

            Rectangle {
                id: iconCaretContainer
                color: getBackgroundColor()
                Layout.minimumWidth: 24
                Layout.minimumHeight: 24
                Layout.preferredWidth: 24
                Layout.preferredHeight: 24
                Layout.alignment: Qt.AlignLeft

                Image {
                    id: iconCaret
                    source: root.state === "folded" ? "../../icons/arrow_right.svg" : "../../icons/arrow_drop_down.svg"
                    antialiasing: true
                    height: 24
                    width: 24
                    sourceSize.height: 24
                    sourceSize.width: 24
                }

                ColorOverlay {
                    anchors.fill: iconCaret
                    source: iconCaret
                    color: Material.color(Material.Grey, Material.Shade700)
                }
            }


            Rectangle {
                id: iconFolderContainer
                color: getBackgroundColor()
                Layout.minimumWidth: 50
                Layout.minimumHeight: 25
                Layout.preferredWidth: 50
                Layout.preferredHeight: 25
                Layout.alignment: Qt.AlignLeft

                Image {
                    id: iconFolder
                    source: "../../icons/folder.svg"
                    antialiasing: true
                    height: 24
                    width: 24
                    sourceSize.height: 24
                    sourceSize.width: 24
                }

                ColorOverlay {
                    anchors.fill: iconFolder
                    source: iconFolder
                    color: Material.color(Material.Grey, Material.Shade700)
                }
            }

            Rectangle {
                id: textContainer
                color: getBackgroundColor()
                Layout.fillWidth: true
                Layout.minimumWidth: 50
                Layout.minimumHeight: 50
                Layout.alignment: Qt.AlignVCenter

                ColumnLayout {
                    Text  {
                        horizontalAlignment:  Text.AlignLeft
                        verticalAlignment: Text.AlignVCenter
                        text: root.model.name
                        clip: true
                        wrapMode: Text.Wrap
                        color: Material.color(Material.Grey, Material.Shade700)
                        font.pixelSize: 14
                    }

                    Text  {
                        text: root.queries ? root.queries.length + " queries" : "0 queries"
                        clip: true
                        elide: Text.ElideRight
                        wrapMode: Text.WordWrap
                        color: Material.color(Material.Grey, Material.Shade700)
                        font.pixelSize: 12
                    }
                }
            }
        }
    }
}
