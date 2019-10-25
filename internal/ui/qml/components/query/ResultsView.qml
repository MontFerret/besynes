import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.13
import "../common" as Common

Item {
    id: root

    property var value: ({ data: "", error: "", stats: { compilation: "", runtime: "", size: "" } })

    Page {
        anchors.fill: parent

        header: Pane {
            height: 30
            width: parent.width
            topPadding: 0
            bottomPadding: 0

            PerfView {
                width: 300
                compile: root.value.stats && root.value.stats.compilation ? root.value.stats.compilation : ""
                runtime: root.value.stats && root.value.stats.runtime ? root.value.stats.runtime : ""
                size: root.value.stats && root.value.stats.size ? root.value.stats.size : ""
            }
        }

        Rectangle {
            anchors.fill: parent
            border.width: 1
            border.color: Material.color(Material.Grey, Material.Shade200)
            radius: 5

            TextEdit {
                id: results
                anchors.fill: parent
                color: value.error ? Material.color(Material.Red) : Material.color(Material.Grey, Material.Shade900)
                focus: true
                text: value.error ? value.error : value.data
                padding: 10
                wrapMode: TextEdit.WordWrap
                readOnly: true
                selectByMouse: true
                selectionColor: Material.color(Material.Purple)
            }
        }
    }
}
