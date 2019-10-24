import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.13
import QtQuick.Controls.Styles 1.4
import QtQuick.Controls.Material 2.13

Item {
    anchors.fill: parent

    function insertContentAt(item, idx) {
        const children = queryContentList.children
        const current = children[idx]
        children.push(item)
        children[idx] = item
        children[children.length -1] = current
    }

    function removeContentAt(idx) {
        const children = queryContentList.children
        const len = children.length
        const new_children = []

        console.log('removing:', idx)

        for (let i = 0; i < len; i++) {
            if (i !== idx) {
                new_children.push(children[i])
            }
        }

        queryContentList.children = new_children
    }

    function newTab() {
        const tabs_count = queryTabList.count
        const idx = tabs_count
        const id = `query_${Math.random().toString()}`

        const btn = tabBtn.createObject(queryTabList, {
            uid: id,
            text: `UNTITLED QUERY ${tabs_count + 1}`
        })
        const content = tabContent.createObject(queryContentList, {
            uid: id
        })

        queryTabList.insertItem(idx, btn)
        insertContentAt(content, idx)
        queryTabList.setCurrentIndex(idx)

        // forcing to re-render the Layout
        // otherwise it fails to scale first rendered item
        queryContentList.width++
        queryContentList.width--
    }

    function closeTab(target_uid) {
        const len = queryTabList.count
        let idx = -1
        let current_item;

        for (var i = 0; i < len; ++i) {
            current_item = queryTabList.itemAt(i)

            if (current_item.uid === target_uid) {
                idx = i
                break
            }
        }

        if (idx > -1) {
            queryTabList.removeItem(current_item)
            removeContentAt(idx)
            queryTabList.setCurrentIndex(idx)
        }
    }

    Page {
        anchors.fill: parent
        header: RowLayout {
            width: parent.width
            spacing: 0

            TabBar {
                id: queryTabList
                Layout.fillWidth: true
                Material.accent: Material.Purple

                Component {
                    id: tabBtn

                    TabButton {
                        property string uid: ''
                        Material.accent: Material.Purple

                        Button {
                            id: tabCloseButton
                            text: "x"
                            width: 25
                            anchors.right: parent.right
                            anchors.top: parent.top
                            anchors.bottom: parent.bottom
                            anchors.rightMargin: 10
                            contentItem: Text {
                                text: tabCloseButton.text
                                font.family: parent.font.family
                                font.bold: false
                                font.pixelSize: 14
                                font.capitalization: Font.AllLowercase
                                horizontalAlignment: Text.AlignHCenter
                                verticalAlignment: Text.AlignVCenter
                            }
                            background: Rectangle {
                                opacity: 0
                                border.width: 0
                            }
                            onClicked: closeTab(this.parent.uid)
                        }
                    }
                }

                Component {
                    id: tabContent
                    QueryEditor {
                        Layout.fillWidth: true
                        Layout.fillHeight: true

                        property string uid: ''
                    }
                }
            }

            TabButton {
                id: tabBtnAdd
                Material.accent: Material.color(Material.Grey, Material.Shade900)
                Layout.alignment: Qt.AlignRight
                text: "+"
                width: 40
                font.bold: false
                font.pixelSize: 20
                hoverEnabled: false
                onClicked: newTab()
            }
        }

        StackLayout {
            id: queryContentList
            anchors.fill: parent
            currentIndex: queryTabList.currentIndex
        }
    }
}
