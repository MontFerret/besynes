import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.13
import "../common" as Common

Item {
    property string compile: ""
    property string runtime: ""
    property string size: ""

    id: root

    RowLayout {
        anchors.fill: parent
        spacing: 12
        visible: root.compile !== "" && root.runtime !== ""

        Common.Statistic {
            id: statCompiletime
            name: "compile"
            value: root.compile
            Layout.fillHeight: true
            Layout.preferredWidth: root.compile ? root.compile.length : 0
        }

        Common.Statistic {
            id: statRuntime
            name: "runtime"
            value: root.runtime
            Layout.fillHeight: true
            Layout.preferredWidth: root.runtime ? root.runtime.length : 0
        }

        Common.Statistic {
            id: statSize
            name: "size"
            value: root.size
            Layout.fillHeight: true
            Layout.preferredWidth: root.size ? root.size.length : 0
        }
    }
}
