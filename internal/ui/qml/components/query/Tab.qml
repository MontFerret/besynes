import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Styles 1.4
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.13
import "../common" as Common

Item {
    property string name: "UNTITLED QUERY"
    property string text: ""
    signal saveResult(string data)

    id: root

    states: [
        State {
            name: "new"
            PropertyChanges { target: execBtn; enabled: true }
            PropertyChanges { target: queryEditor; enabled: true }
            PropertyChanges { target: paramsEditor; enabled: true }
            PropertyChanges { target: resultsViewer; enabled: true }
            PropertyChanges { target: progress; visible: false }
        },
        State {
            name: "ready"
            PropertyChanges { target: execBtn; enabled: true }
            PropertyChanges { target: queryEditor; enabled: true }
            PropertyChanges { target: paramsEditor; enabled: true }
            PropertyChanges { target: resultsViewer; enabled: true }
            PropertyChanges { target: progress; visible: false }
        },
        State {
            name: "loading"
            PropertyChanges { target: execBtn; enabled: false }
            PropertyChanges { target: queryEditor; enabled: false }
            PropertyChanges { target: paramsEditor; enabled: false }
            PropertyChanges { target: resultsViewer; enabled: false }
            PropertyChanges { target: progress; visible: true }
        }
    ]

    Component.onCompleted: {
        root.state = "new"
    }

    QtObject {
        id: query
        property string text: ""
        property string params: ""
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

                ProgressBar {
                    id: progress
                    Material.accent: Material.color(Material.Indigo, Material.Shade800)
                    anchors.bottom: parent.bottom
                    width: parent.width
                    indeterminate: true
                }
            }

            RowLayout {
                anchors.fill: parent

                Button {
                    Material.background: Material.Indigo
                    Layout.alignment: Qt.AlignRight
                    Layout.bottomMargin: 5
                    id: execBtn
                    text: "Run"
                    highlighted: true
                    onClicked: {
                        // if true, we in QtCreator
                        if (typeof queryApi === "undefined") {
                            root.state = "loading"
                            resultsViewer.value = {
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
                                text: query.text,
                                params: query.params,
                            }, (result) => {
                                resultsViewer.value = result

                                root.state = "ready"
                            })
                        } catch (e) {
                            resultsViewer.value = {
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
                                onEditingFinished: text => query.text = text
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
                                onEditingFinished: text => query.params = text
                            }
                        }
                    }
                }

                Pane {
                    id: resultsPane
                    SplitView.maximumHeight: pageContent.height
                    SplitView.preferredHeight: root.state !== "new" ? pageContent.height / 2 : 0
                    SplitView.minimumHeight: 0
                    padding: 5
                    leftPadding: 12
                    rightPadding: 12
                    bottomPadding: 12

                    ResultsViewer {
                        id: resultsViewer
                        anchors.fill: parent
                        onSave: (data) => {
                            if(data && root.saveResult) {
                                root.saveResult(data)
                            }
                        }
                    }
                }
            }
        }
    }
}
