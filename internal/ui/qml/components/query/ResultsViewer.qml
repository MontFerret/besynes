import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.13
import "../common" as Common

Item {
    id: root

    property var value: ({ data: "", error: "", stats: { compilation: "", runtime: "", size: "" } })
    signal save(string text)

    Page {
        anchors.fill: parent

        header: Pane {
            padding: 0
            bottomPadding: 15

            RowLayout {
                anchors.fill: parent
                spacing: 0

                PerfViewer {
                    id: perfStats
                    Layout.alignment: Qt.AlignLeft
                    width: 300
                    compile: root.value.stats && root.value.stats.compilation ? root.value.stats.compilation : ""
                    runtime: root.value.stats && root.value.stats.runtime ? root.value.stats.runtime : ""
                    size: root.value.stats && root.value.stats.size ? root.value.stats.size : ""
                }

                Item {
                    Layout.alignment: Qt.AlignRight
                    Layout.fillWidth: true
                    Layout.preferredWidth: copyBtn.width + saveBtn.width
                    Layout.fillHeight: true
                    Layout.preferredHeight: copyBtn.height

                    TabButton {
                        id: copyBtn
                        anchors.right: saveBtn.left
                        Material.foreground: Material.color(Material.Grey, Material.Shade700)
                        Material.accent: Material.color(Material.Grey, Material.Shade700)
                        icon.width: 20
                        icon.height: 20
                        antialiasing: true
                        icon.source: `../../icons/copy.svg`
                        onClicked: {
                            viewer.copy()
                        }
                    }

                    TabButton {
                        id: saveBtn
                        anchors.right: parent.right
                        Material.foreground: Material.color(Material.Grey, Material.Shade700)
                        Material.accent: Material.color(Material.Grey, Material.Shade700)
                        icon.width: 20
                        icon.height: 20
                        antialiasing: true
                        icon.source: `../../icons/save.svg`
                        onClicked: {
                            if (root.save) {
                                root.save(viewer.text)
                            }
                        }
                    }
                }
            }
        }

        Common.CodeEditor {
            id: viewer
            anchors.fill: parent
            text: value.error ? value.error : value.data
            color: value.error ? Material.color(Material.Red) : Material.color(Material.Grey, Material.Shade900)
            readOnly: true
        }
    }
}
