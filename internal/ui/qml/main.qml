import QtQuick 2.13
import Qt.labs.platform 1.1
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13
import QtQuick.Layouts 1.12
import Qt.labs.platform 1.1 as Labs
import "./components/common" as Common
import "./components/query" as Query
import "./components/settings" as Settings
import "./components/catalog" as Catalog

ApplicationWindow {
    id: win
    visible: true
    width: 1024
    height: 768
    minimumWidth: 512
    minimumHeight: 384
    title: "Besynes"

    header: ToolBar {
        Material.background: Material.DeepPurple
        leftPadding: 15
        rightPadding: 15

        RowLayout {
            anchors.fill: parent

            RoundButton {
                Layout.alignment: Qt.AlignLeft
                icon.source: "./icons/menu.svg"
                flat: true
                onClicked: {
                    drawer.open()
                }
            }

            Rectangle {
                Layout.fillHeight: true
                Layout.fillWidth: true
                Layout.alignment: Qt.AlignCenter
                color: "transparent"

                Image {
                    anchors.centerIn: parent
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

            RoundButton {
                Layout.alignment: Qt.AlignRight
                icon.source: "./icons/settings.svg"
                flat: true
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
        onError: (err) => {
            alert.open({ type: 'error', title: "Error", body: err });
        }
    }

    Common.Alert {
        id: alert
    }

    Drawer {
        id: drawer
        width: 0.4 * parent.width
        height: parent.height
        visible: false

        Catalog.View {
            id: catalog
            anchors.fill: parent
            onSelected: (groupId, queryId) => {
                tabView.addQuery(groupId, queryId);
                drawer.close();
            }
        }
    }

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
