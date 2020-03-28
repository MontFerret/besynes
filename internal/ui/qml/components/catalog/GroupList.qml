import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.13
import "../common" as Common

Item {
    id: root
    property ListModel model;
    signal selected(var group);

    ListView {
        anchors.fill: parent
        spacing: 0
        model: root.model
        delegate: Component {
            GroupListItem {
                width: parent.width
                height: 80
                model: root.model.get(index)
                onSelected: {
                    const group = collectionModel.get(index);

                    if (root.selected) {
                        root.selected(group);
                    }
                }
            }
        }
    }
}
