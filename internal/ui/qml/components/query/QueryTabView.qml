import QtQuick 2.13
import QtQuick.Controls 1.4 as C1
import QtQuick.Controls 2.13
import QtQuick.Controls.Styles 1.4
import QtQuick.Controls.Material 2.13

Item {
    anchors.fill: parent

    Component {
        id: editor
        QueryEditor {}
    }

    C1.TabView {
        id: queryTabsList
        anchors.fill: parent
        style: TabViewStyle {
            tab: Rectangle {
                function isButton() {
                    return !styleData || styleData.title === '+';
                }

                implicitWidth: !isButton() ? 175 : 40
                implicitHeight: 40
                color: Material.color(Material.Grey, Material.Shade200)
                anchors.right: parent.right
                anchors.top: parent.top

                TabButton {
                    id: tabButton
                    anchors.fill: parent
                    contentItem: Text {
                        anchors.fill: parent
                        text: styleData.title
                        font.family: parent.font.family
                        font.bold: false
                        font.capitalization: Font.AllUppercase
                        horizontalAlignment: Text.AlignHCenter
                        verticalAlignment: Text.AlignVCenter
                        elide: Text.ElideRight
                    }
                    onClicked: {
                        if (isButton()) {
                            var tab_count = queryTabsList.count
                            var t = queryTabsList.insertTab(tab_count > 0 ? tab_count - 1 : 0, "Untitled Query", editor)
                            t.active = true;

                            queryTabsList.currentIndex = tab_count - 1
                        } else {
                            queryTabsList.currentIndex = styleData.index
                        }
                    }

                    // Close button
                    Button {
                        id: tabCloseButton
                        text: !isButton() ? "x" : ""
                        visible: !isButton()
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

                    // Bottom line of a selected tab
                    Rectangle {
                        width: parent.width
                        height: 3
                        anchors.bottom: parent.bottom
                        color: "#81D4FA"
                        visible: !isButton() && styleData.selected
                    }
                }
            }

            frame: styleData ? editor : null
        }

        C1.Tab {
            id: tabAdd
            title: qsTr("+")
        }
    }
}
