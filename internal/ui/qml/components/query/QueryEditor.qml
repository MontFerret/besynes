import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.13
import besynes.models.query 1.0

Item {
    property string name: "UNTITLED QUERY"
    property string text: ""

    id: root
    anchors.fill: parent

    states: [
        State {
            name: "ready"
            PropertyChanges { target: execBtn; enabled: true }
            PropertyChanges { target: queryEditor; enabled: true }
        },
        State {
            name: "loading"
            PropertyChanges { target: execBtn; enabled: false }
            PropertyChanges { target: queryEditor; enabled: false }
        }
    ]

    Component.onCompleted: {
        root.state = "ready"
    }

    Page {
        anchors.fill: parent
        anchors.topMargin: 2
        padding: 5

        header: ToolBar {
            Material.background: Material.DeepPurple
            leftPadding: 15
            rightPadding: 15

            RowLayout {
                anchors.fill: parent

                Label {
                    text: name
                    elide: Label.ElideRight
                    horizontalAlignment: Qt.AlignHCenter
                    verticalAlignment: Qt.AlignVCenter
                    Layout.fillWidth: true
                }

                Button {
                    Material.background: Material.Blue
                    id: execBtn
                    text: "Exec"
                    highlighted: true
                    onClicked: {
                        root.state = "loading"

                        try {
                            // const result
                            resultsView.value = execution.execute(queryEditor.text, qsTr(""))
                        } catch (e) {
                            resultsView.value = {
                                data: "",
                                error: e.toString()
                            }
                        } finally {
                            root.state = "ready"
                        }
                    }
                }
            }
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
