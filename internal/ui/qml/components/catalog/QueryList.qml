import QtQuick 2.0

Item {
    id: root
    property var model: ({ id: "", name: "", queries: [] })
    signal selected(string groupId, string queryId)
    signal edited(string groupId, string queryId)
    signal deleted(string groupId, string queryId)

    ListView {
        anchors.fill: parent
        spacing: 0
        model: root.model.queries
        delegate: Component {
            QueryListItem {
                width: parent.width
                height: 80
                model: root.model.queries.get(index)
                onSelected: (i) => root.selected(root.model.id, i)
                onEdited: (i) => root.edited(root.model.id, i)
                onDeleted: (i) => root.deleted(root.model.id, i)
            }
        }
    }
}
