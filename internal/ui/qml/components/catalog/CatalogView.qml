import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.13
import "../common" as Common

Control {
    id: root

    ListModel {
        id: collectionModel
    }

    QtObject {
        id: selection
        property int groupIndex: -1
    }

    state: "catalog"
    states: [
        State {
            name: "catalog"; when: views.depth === 1
            PropertyChanges { target: title; text: "Catalog" }
            PropertyChanges { target: backBtn; x: -50; y: 0 }
        },

        State {
            name: "group"; when: views.depth === 2
            PropertyChanges { target: title; text: collectionModel.get(selection.groupIndex).name }
            PropertyChanges { target: backBtn; x: 0; y: 0; }
        }
    ]

    Component.onCompleted: {
        const qd = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer imperdiet libero in massa pulvinar imperdiet.";
        const gd = "Pellentesque tempus molestie eleifend. Integer ex elit, laoreet et diam at, venenatis sagittis dui. Vestibulum tincidunt urna nec lacus molestie, non tincidunt sem gravida.";

        for (var i = 0; i <= 100; i ++) {
            const queries = [];

            for (var y = 0; y <= 10; y ++) {
                queries.push({
                    id: y + 1,
                    name: "Query " + (y + 1),
                    description: qd
                })
            }

            collectionModel.append({
                id: i + 1,
                name: "Group " + (i + 1),
                description: gd,
                queries: queries
            })
        }
    }

    Component {
        id: groupsView

        ListView {
            spacing: 0
            model: collectionModel
            delegate: Component {
                CatalogGroup {
                    width: parent.width
                    height: 80
                    model: collectionModel.get(index)
                    onSelected: {
                        selection.groupIndex = index

                        const group = collectionModel.get(index);

                        views.push(itemsView, {
                            model: group.queries || []
                        })
                    }
                }
            }
        }
    }

    Component {
        id: itemsView

        ListView {
            id: itemsViewInstance
            spacing: 0
            delegate: Component {
                CatalogGroupItem {
                    width: parent.width
                    height: 80
                    model: itemsViewInstance.model.get(index)
                }
            }
        }
    }

    Page {
        anchors.fill: parent
        header: ToolBar {
            Material.background: Material.DeepPurple
            leftPadding: 15
            rightPadding: 15

            RoundButton {
                id: backBtn
                x: -50
                icon.source: "../../icons/arrow_back-black.svg"
                flat: true
                onClicked: views.pop()

                Behavior on x {
                    NumberAnimation { properties: "x,y"; easing.type: Easing.InOutQuad; duration: 200; loops: 1 }
                }
            }

            RowLayout {
                anchors.fill: parent

                Rectangle {
                    Layout.fillWidth: true
                    Layout.fillHeight: true
                    Layout.alignment: Qt.AlignLeft | Qt.AlignVCenter
                    color: "transparent"

                    Text {
                        id: title
                        anchors.centerIn: parent
                        color: "white"
                        font.pixelSize: 16
                        font.family: "Roboto"
                        font.weight: Font.ExtraBold
                        antialiasing: true
                        text: "Catalog"
                    }
                }

                RoundButton {
                    Layout.alignment: Qt.AlignRight
                    Layout.minimumWidth: 50
                    Layout.minimumHeight: 50
                    Layout.preferredWidth: 50
                    Layout.preferredHeight: 50
                    icon.source: "../../icons/search-black.svg"
                    flat: true
                }
            }
        }

        StackView {
            id: views
            initialItem: groupsView
            anchors.fill: parent
        }
    }
}
