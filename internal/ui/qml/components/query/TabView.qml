import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.13
import QtQuick.Controls.Styles 1.4
import QtQuick.Controls.Material 2.13

Control {
    signal saveResult(string query, string data)

    id: root
    anchors.fill: parent

    QtObject {
        id: state
        property int counter: 0
    }

    function addQuery(groupId, queryId) {
        if (typeof catalogApi === "undefined") {
            newTab()
            return
        }

        catalogApi.getQuery(queryId, (err, query) => {
            if (err) {
                console.error(err);
                newTab();
                return;
            }

            addTab(query);
        });
    }

    function addTab(query) {
        state.counter += 1

        tabListModel.append(query)

        // forcing to re-render the Layout
        // otherwise it fails to scale first rendered item
        queryContentList.width++
        queryContentList.width--

        queryTabBar.setCurrentIndex(tabListModel.count - 1)
    }

    function newTab() {
        const uid = `query_${state.counter}`
        const name = `UNTITLED QUERY ${state.counter + 1}`

        addTab({ uid, name });
    }

    function closeTab(target_uid) {
        if (target_uid < 0) {
           return;
        }

        const len = tabListModel.count
        let idx = -1

        for (let i = 0; i < len; i += 1) {
            if (tabListModel.get(i).uid === target_uid) {
                idx = i
                break
            }
        }

        if (idx > -1) {
            tabListModel.remove(idx)
        }
    }

    ListModel {
        id: tabListModel
    }

    Page {
        anchors.fill: parent
        header: RowLayout {
            width: parent.width
            spacing: 0

            TabBar {
                id: queryTabBar
                Layout.fillWidth: true
                Material.accent: Material.Purple

                Repeater {
                    model: tabListModel
                    delegate: TabButton {
                        text: name
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
                            onClicked: closeTab(uid)
                        }
                    }
                }
            }

            RoundButton {
                id: tabBtnAdd
                Material.foreground: Material.color(Material.Grey, Material.Shade700)
                Material.accent: Material.color(Material.Grey, Material.Shade700)
                Layout.alignment: Qt.AlignRight
                icon.source: "../../icons/add.svg"
                icon.width: 24
                icon.height: 24
                width: 50
                flat: true
                hoverEnabled: false
                onClicked: newTab()
            }
        }

        StackLayout {
            id: queryContentList
            anchors.fill: parent
            currentIndex: queryTabBar.currentIndex

            Repeater {
                model: tabListModel
                delegate: Tab {
                    name: name
                    text: text
                    Layout.fillWidth: true
                    Layout.fillHeight: true
                    onSaveResult: (data) => {
                        if (root.saveResult && data) {
                            root.saveResult(this.name, data)
                        }
                    }
                }
            }
        }

        Text {
            anchors.centerIn: parent
            visible:tabListModel.count == 0
            text: `Click "+" icon\nor\n${Qt.platform.os !== 'osx' ? "Ctrl" : "Cmd"}+N to create a new query`
            horizontalAlignment: Text.AlignHCenter
            color: Material.color(Material.Grey, Material.Shade700)
            font.family: "Roboto"
            font.pixelSize: 16
        }

        Shortcut {
            sequence: "Ctrl+N"
            onActivated: newTab()
        }

        Shortcut {
            sequence: "Ctrl+W"
            onActivated: {
                const idx = queryContentList.currentIndex

                if (idx < 0) {
                    return
                }

                const tab = tabListModel.get(idx)

                closeTab(tab.uid)
            }
        }
    }
}
