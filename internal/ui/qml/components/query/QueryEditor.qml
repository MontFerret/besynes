import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.13

Item {
    property string name: "UNTITLED QUERY"
    property string text: ""

    id: root

    states: [
        State {
            name: "ready"
            PropertyChanges { target: execBtn; enabled: true }
            PropertyChanges { target: queryEditor; enabled: true }
            PropertyChanges { target: spinner; running: false }
        },
        State {
            name: "loading"
            PropertyChanges { target: execBtn; enabled: false }
            PropertyChanges { target: queryEditor; enabled: false }
            PropertyChanges { target: spinner; running: true }
        }
    ]

    Component.onCompleted: {
        root.state = "ready"
    }

    Page {
        id: page
        anchors.fill: parent
        anchors.topMargin: 2
        padding: 5

        header: ToolBar {
            leftPadding: 15
            rightPadding: 15
            background: Rectangle {
                anchors.fill: parent
                color: Material.color(Material.Grey, Material.Shade50)

                Rectangle {
                    width: parent.width
                    height: 1
                    anchors.bottom: parent.bottom
                    color: Material.color(Material.Grey, Material.Shade200)
                }
            }

            RowLayout {
                anchors.fill: parent

                Button {
                    Material.background: Material.Blue
                    Layout.alignment: Qt.AlignRight
                    Layout.bottomMargin: 5
                    Layout.rightMargin: 5
                    id: execBtn
                    text: "Run"
                    highlighted: true
                    onClicked: {
                        root.state = "loading"

                        try {
                            // const result
                            queryApi.execute({
                                text: queryEditor.text
                            }, (err, result) => {
                                resultsView.value = {
                                    data: result,
                                    error: err
                                }

                                root.state = "ready"
                            })
                        } catch (e) {
                            resultsView.value = {
                                data: "",
                                error: e.toString()
                            }

                            root.state = "ready"
                        }
                    }
                }
            }
        }

        Rectangle {
            anchors.fill: parent

            BusyIndicator {
                id: spinner
                anchors.centerIn: parent
                running: false
                z: 90
            }

            SplitView {
                anchors.fill: parent
                orientation: Qt.Vertical

                Rectangle {
                    SplitView.fillWidth: true
                    SplitView.fillHeight: true
                    SplitView.minimumHeight: 200

                    SplitView {
                        anchors.fill: parent
                        orientation: Qt.Horizontal

                        Pane {
                            SplitView.fillWidth: true
                            SplitView.minimumWidth: 200
                            id: queryPane

                            CodeEditor {
                                id: queryEditor
                                anchors.fill: parent
                                text: text
                            }
                        }

                        Pane {
                            SplitView.minimumWidth: 150

                            ParamsEditor {
                                anchors.fill: parent
                                id: paramsEditor
                            }
                        }
                    }
                }

                Pane {
                    SplitView.minimumHeight: 150

                    ResultsView {
                        id: resultsView
                        anchors.fill: parent
                    }
                }
            }
        }
    }
}
