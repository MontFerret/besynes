import QtQuick 2.13
import QtQuick.Controls 1.4
import QtQuick.Controls 2.5
import QtQuick.Controls.Styles 1.4
import QtQuick.Controls.Material 2.12

Item {
    anchors.fill: parent

    Component {
        id: editor
        QueryEditor {}
    }

    TabView {
        id: queryTabsList
        anchors.fill: parent
        style: TabViewStyle {
            tab: Rectangle {
                function isButton() {
                    return !styleData || styleData.title === '+';
                }

                implicitWidth: !isButton() ? 175 : 40
                implicitHeight: 40

                TabButton {
                    id: tabButton
                    anchors.fill: parent
                    text: styleData.title
                    onClicked: {
                        if (tabButton.text === "+") {
                            var tab_count = queryTabsList.count
                            var t = queryTabsList.insertTab(tab_count > 0 ? tab_count - 1 : 0, "Untitled Query", editor)
                            t.active = true; // real loading

                            queryTabsList.currentIndex = tab_count - 1
                        } else {
                            queryTabsList.currentIndex = styleData.index
                        }
                    }

                    Button {
                        id: tabCloseButton
                        text: tabButton.text !== "+" ? "x" : ""
                        visible: false
                        width: 25
                        anchors.right: parent.right
                        contentItem: Text {
                            anchors.fill: parent
                            text: tabCloseButton.text
                            font.family: parent.font.family
                            font.bold: true
                            font.capitalization: Font.AllLowercase
                            horizontalAlignment: Text.AlignHCenter
                            verticalAlignment: Text.AlignVCenter
                            elide: Text.ElideRight
                        }
                        background: Rectangle {
                            opacity: 0
                            border.width: 0
                        }
                        onClicked: {
                            queryTabsList.removeTab(styleData.index)
                        }
                    }

                    Rectangle {
                        width: parent.width
                        height: 3
                        anchors.bottom: parent.bottom
                        color: "#81D4FA"
                        visible: styleData.title !== '+' && styleData.selected
                    }

//                    MouseArea {
//                        anchors.fill: parent
//                        preventStealing: true
//                        enabled: tabButton.text !== "+"
//                        hoverEnabled: tabButton.text !== "+"
//                        onEntered: {
//                            if (tabButton.text !== "+") {
//                                tabCloseButton.visible = true
//                            }
//                        }
//                        onExited: {
//                            tabCloseButton.visible = false
//                        }
//                    }
                }
            }
            frame: styleData && styleData.title !== '+' ? editor : null
        }

        Tab {
            id: tabAdd
            title: qsTr("+")
        }
    }
}
