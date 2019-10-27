import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13
import QtQuick.Layouts 1.12
import Qt.labs.platform 1.1
import "./components/query" as Query

ApplicationWindow {
    id: win
    visible: true
    width: 1024
    height: 768
    title: "Besynes"

    header: ToolBar {
        Material.background: Material.DeepPurple
        leftPadding: 15
        rightPadding: 15

        RowLayout {
            anchors.fill: parent

            Label {
                text: "BESYNES"
                elide: Label.ElideRight
                horizontalAlignment: Qt.AlignLeft
                verticalAlignment: Qt.AlignVCenter
                Layout.fillWidth: true
            }
        }
    }

    background: Rectangle {
        color: Material.color(Material.Grey, Material.Shade200)
        anchors.fill: parent
    }

    FileDialog {
        id: fileDialog
        title: "Please choose a file"
        fileMode: FileDialog.SaveFile
        folder: StandardPaths.writableLocation(StandardPaths.DocumentsLocation)
    }

    Query.TabView {
        id: tabs
        onSaveResult: (query, data) => {
            const fileName = `${query}.json`
            fileDialog.title = "Save query results"
            fileDialog.currentFile = fileName
            fileDialog.open()
        }
    }
}
