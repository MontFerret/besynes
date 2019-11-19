import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13
import QtQuick.Layouts 1.12
import Qt.labs.platform 1.1 as Labs
import "./components/common" as Common
import "./components/query" as Query
import "./components/settings" as Settings
import "./components/catalog"

ApplicationWindow {
    id: win
    visible: true
    width: 1024
    height: 768
    title: "Besynes"

    QtObject {
        id: settingsModel

        property string cdpAddress: "http://127.0.0.1:9222"
    }

    header: ToolBar {
        Material.background: Material.DeepPurple
        leftPadding: 15
        rightPadding: 15

        RowLayout {
            anchors.fill: parent

            Rectangle {
                Layout.alignment: Qt.AlignLeft

                Image {
                    id: logoImg
                    source: "./images/logo.png"
                    width: 50
                    height: 50
                    antialiasing: true

                    RotationAnimator {
                        id: logoImgAnimation
                        target: logoImg;
                        from: 0;
                        to: 360;
                        duration: 500
                        running: false
                    }

                    MouseArea {
                        anchors.fill: parent
                        onDoubleClicked: {
                            logoImgAnimation.running = true
                        }
                    }
                }
            }

            TabButton {
                Layout.alignment: Qt.AlignRight
                icon.source: "./icons/settings.svg"
                Material.foreground: "white"
                Material.accent: "white"
                onClicked: {
                    settingsDialog.open()
                }
            }
        }
    }

    background: Rectangle {
        color: Material.color(Material.Grey, Material.Shade200)
        anchors.fill: parent
    }

    Labs.FileDialog {
        id: fileDialog
        title: "Please choose a file"
        fileMode: Labs.FileDialog.SaveFile
        folder: Labs.StandardPaths.writableLocation(Labs.StandardPaths.DocumentsLocation)
    }

    Settings.Dialog {
        id: settingsDialog
    }

    Component {
        id: splitHandle

        Common.SplitHandle {}
    }

    SplitView {
        anchors.fill: parent
        orientation: Qt.Horizontal
        handle: splitHandle

        Rectangle {
            SplitView.maximumWidth: win.width / 2
            SplitView.preferredWidth: 0
            SplitView.minimumWidth: 0

            CatalogView {
                id: catalogView
                anchors.fill: parent
            }
        }

        Rectangle {
            SplitView.fillWidth: true
            SplitView.fillHeight: true
            SplitView.minimumWidth: 0
            SplitView.preferredWidth: 200

            Query.TabView {
                id: tabView
                anchors.fill: parent
                onSaveResult: (query, data) => {
                    const fileName = `${query}.json`
                    fileDialog.title = "Save query results"
                    fileDialog.currentFile = fileName
                    fileDialog.open()
                }
            }
        }
    }
}
