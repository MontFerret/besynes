import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13
import "../common" as Common

Control {
    id: root

    ListModel {
        id: collectionModel
    }

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

    Rectangle {
        color: Material.color(Material.Grey, Material.Shade100)
        anchors.fill: parent

        Component {
            id: listItemDelegate

            CatalogGroup {
                width: listView.width
                height: 80
                model: collectionModel.get(index)
            }
        }

        ListView {
            id: listView
            anchors.fill: parent
            model: collectionModel
            delegate: listItemDelegate
            snapMode: ListView.SnapToItem
            headerPositioning: ListView.OverlayHeader
            header: Pane {
                width: parent.width
                z: 2
                background: Rectangle {
                    color: Material.color(Material.Grey, Material.Shade200)
                    anchors.fill: parent
                    border.color: Material.color(Material.Grey, Material.Shade300)
                }
            }
        }
    }
}
