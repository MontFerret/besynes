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
                        Material.foreground: Material.color(Material.Grey, Material.Shade700)
                        Material.accent: Material.Purple
                        font.hintingPreference: Font.PreferFullHinting
                        onHoveredChanged: tabCloseButton.visible = hovered

                        Button {
                            id: tabCloseButton
                            anchors.right: parent.right
                            anchors.top: parent.top
                            anchors.bottom: parent.bottom
                            width: 50
                            height: 50
                            visible: false
                            icon.height: 18
                            icon.width: 18
                            icon.source: "../../icons/clear.svg"
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
                    QueryTab {
                        Layout.fillWidth: true
                        Layout.fillHeight: true

                        property string uid: ''
                    }
                }
            }

            TabButton {
                id: tabBtnAdd
                Material.foreground: Material.color(Material.Grey, Material.Shade700)
                Material.accent: Material.color(Material.Grey, Material.Shade700)
                Layout.alignment: Qt.AlignRight
                icon.source: "../../icons/add.svg"
                icon.width: 24
                icon.height: 24
                width: 50
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
