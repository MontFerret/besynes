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
            name: "new"
            PropertyChanges { target: execBtn; enabled: true }
            PropertyChanges { target: queryEditor; enabled: true }
            PropertyChanges { target: spinner; running: false }
        },
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
        root.state = "new"
    }

    QtObject {
        id: query
        property string text: ""
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
                    id: execBtn
                    text: "Run"
                    highlighted: true
                    onClicked: {
                        // if true, we in QtCreator
                        if (typeof queryApi === "undefined") {
                            root.state = "loading"
                            resultsView.value = {
                                data: query.text,
                                error: ""
                            }
                            root.state = "ready"

                            return
                        }

                        root.state = "loading"

                        try {
                            // const result
                            queryApi.execute({
                                text: query.text
                            }, (result) => {
                                resultsView.value = result

                                root.state = "ready"
                            })
                        } catch (e) {
                            resultsView.value = {
                                data: "",
                                error: e.toString(),
                                stats: {}
                            }

                            root.state = "ready"
                        }
                    }
                }
            }
        }

        Rectangle {
            id: pageContent
            anchors.fill: parent

            BusyIndicator {
                id: spinner
                anchors.centerIn: parent
                running: false
                z: 90
            }

            Component {
                id: splitHandle

                Rectangle {
                    function isVertical() {
                        return parent.orientation === Qt.Vertical
                    }

                    id: root
                    implicitWidth: 8
                    implicitHeight: 8
                    color: Material.color(Material.Grey, Material.Shade50)
                    state: SplitHandle.pressed ? "pressed" : "released"
                    states: [
                        State {
                            name: "released"
                            PropertyChanges {
                                target: handle;
                                width: isVertical() ? 10 : 1;
                                height: isVertical() ? 1 : 10;
                                radius: 0
                            }
                        },

                        State {
                            name: "pressed"
                            PropertyChanges {
                                target: handle;
                                width: 3;
                                height: 3;
                                radius: 3
                            }
                        }
                    ]

                    Rectangle {
                        id: handle
                        anchors.centerIn: parent
                        color: Material.color(Material.Grey)

                        Behavior on height {
                            PropertyAnimation {
                                easing.type: Easing.InQuad;
                                duration: 100
                            }
                        }

                        Behavior on width {
                            PropertyAnimation {
                                easing.type: Easing.InQuad;
                                duration: 100
                            }
                        }
                    }
                }
            }

            SplitView {
                anchors.fill: parent
                orientation: Qt.Vertical
                handle: splitHandle

                Rectangle {
                    SplitView.fillWidth: true
                    SplitView.fillHeight: true
                    SplitView.minimumHeight: 0
                    SplitView.preferredHeight: 200

                    SplitView {
                        anchors.fill: parent
                        orientation: Qt.Horizontal
                        handle: splitHandle

                        Pane {
                            id: queryPane
                            SplitView.fillWidth: true
                            SplitView.minimumWidth: 0
                            SplitView.preferredWidth: 200
                            padding: 5
                            leftPadding: 12

                            QueryEditor {
                                id: queryEditor
                                anchors.fill: parent
                                text: text
                                onEditingFinished: (text) => {
                                    query.text = text
                                }
                            }
                        }

                        Pane {
                            SplitView.maximumWidth: pageContent.width / 2
                            SplitView.preferredWidth: 0
                            SplitView.minimumWidth: 0
                            padding: 5
                            rightPadding: 12

                            ParamsEditor {
                                anchors.fill: parent
                                id: paramsEditor
                            }
                        }
                    }
                }

                Pane {
                    id: resultsPane
                    SplitView.maximumHeight: pageContent.height / 2
                    SplitView.preferredHeight: root.state !== "new" ? pageContent.height / 2 : 0
                    SplitView.minimumHeight: 0
                    padding: 5
                    leftPadding: 12
                    rightPadding: 12
                    bottomPadding: 12

                    ResultsViewer {
                        id: resultsView
                        anchors.fill: parent
                    }
                }
            }
        }
    }
}
