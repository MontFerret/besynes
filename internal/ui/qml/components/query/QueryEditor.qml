import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.13

Item {
    property string name: qsTr("UNTITLED QUERY")
    property string text: qsTr("")

    anchors.fill: parent

    Page {
        anchors.fill: parent
        anchors.topMargin: 2
        padding: 15

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
                    id: execBtn
                    text: "Exec"
                    highlighted: true
                    Material.background: Material.Blue
                }
            }
        }

        Rectangle {
            anchors.fill: parent

            SplitView {
                anchors.fill: parent

                Rectangle {
                    id: queryPane
                    SplitView.fillWidth: true
                    SplitView.minimumWidth: 200

                    ColumnLayout {
                        spacing: 2

                        Text {
                            id: queryTitle
                            text: "query"
                            opacity: 0.8
                            padding: 10
                        }

                        Rectangle {
                            Layout.fillWidth: true
                            border.width: 1
                            border.color: "#EEEEEE"
                            radius: 5

                            TextEdit {
                                id: query
                                anchors.fill: parent
                                color: "black"
                                padding: 10
                                focus: true
                                text: text
                            }
                        }
                    }
                }

                Rectangle {
                    // implicitWidth: parent.width / 4
                    SplitView.minimumWidth: 150
                    // SplitView.preferredWidth: parent.width / 2

                    ColumnLayout {
                        spacing: 2

                        Text {
                            text: "params"
                            opacity: 0.8
                            padding: 10
                        }

                        Rectangle {
//                            anchors.fill: parent
                            border.width: 1
                            border.color: "#EEEEEE"
                            radius: 5

                            TextEdit {
                                id: params
                                anchors.fill: parent
                                color: "black"
                                padding: 10
                                focus: true
                                text: text
                            }
                        }
                    }
                }
            }




//            Grid {
//                columns: 1
//                spacing: 0
//                anchors.fill: parent

//                Rectangle {
//                    border.color: "#EEEEEE"
//                    border.width: 1
//                    width: parent.width
//                    height: 40

//                    Text {
//                        text: "Result"
//                        color: "#EEEEEE"
//                        opacity: 0.5
//                    }
//                }

//                TextEdit {
//                    id: params
//                    // anchors.fill: parent
//                    color: "black"
//                    focus: true
//                    text: text
//                }
//            }
        }
    }
}
