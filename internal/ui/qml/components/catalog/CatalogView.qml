import QtQuick 2.13
import QtQuick.Controls 2.13
import "../common" as Common

Control {
    id: root


    ListModel {
        id: collectionModel

        ListElement {
            name: "Fegrase"
        }

        ListElement {
            name: "Ferret"
        }
    }

    Component {
        id: listItemDelegate
        Row {
            spacing: 10
            Text { text: name }
        }
    }


    ListView {
        id: listView
        anchors.fill: parent
        model: collectionModel
        delegate: listItemDelegate
        header: Pane {
            width: parent.width
        }
    }

    Common.Line {
        orientation: Qt.Vertical
        align: Qt.AlignTop
    }
}
