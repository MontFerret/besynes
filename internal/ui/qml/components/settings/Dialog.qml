import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13
import QtQuick.Layouts 1.12
import "../common" as Common

Item {
    readonly property var open: () => {
        dialog.open()
    }

    signal error(string text)

    id: root
    width: 400
    height: 200
    anchors.centerIn: parent

    states: [
        State {
            name: "current"
            PropertyChanges { target: dialog; loading: false }
            PropertyChanges { target: generalSettingsForm; enabled: true }
            PropertyChanges { target: cancelBtn; enabled: true }
            PropertyChanges { target: saveBtn; enabled: false }
        },
        State {
            name: "stale"
            PropertyChanges { target: dialog; loading: false }
            PropertyChanges { target: generalSettingsForm; enabled: true }
            PropertyChanges { target: cancelBtn; enabled: true }
            PropertyChanges { target: saveBtn; enabled: true }
        },
        State {
            name: "loading"
            PropertyChanges { target: dialog; loading: true }
            PropertyChanges { target: generalSettingsForm; enabled: false }
            PropertyChanges { target: cancelBtn; enabled: false }
            PropertyChanges { target: saveBtn; enabled: false }
        }
    ]

    QtObject {
        id: settings
        property string cdpAddress: ""
    }

    function noop() {}

    function handler(err) {
        if (err && root.error) {
            if (typeof err === "string") {
                root.error(err);
            } else {
                root.error(err.toString());
            }
        }
    }

    function loadData(cb) {
        if (typeof settingsApi === 'undefined') {
            cb()
            return;
        }

        settingsApi.get((err, values) => {
            if (err) {
                cb(err)
                return
            }

            // TODO: Update to JSON case keys
            cb(JSON.stringify(values))
            settings.cdpAddress = values.settings.cpdAddress
        })
    }

    function saveData(cb) {
        if (typeof settingsApi === 'undefined') {
            root.state = "current"
            cb()
            return
        }

        root.state = "loading"

        settingsApi.save({
            cdpAddress: settings.cdpAddress
        }, (err) => {
            if (err) {
                cb(err)
                root.state = "stale"
                return
            }

            cb("Saved")
            root.state = "current"
        })
    }

    Component.onCompleted: {
        root.state = "current"

        loadData(handler)
    }

    Common.Dialog {
        id: dialog
        title: "Settings"
        width: root.width
        height: root.height
        padding: 15

        GridLayout {
            anchors.fill: parent
            columns: 1

            GeneralForm {
                id: generalSettingsForm
                Layout.fillHeight: true
                Layout.fillWidth: true
                cdpAddress: settings.cdpAddress
                onChanged: (text) => {
                    root.state = "stale"
                    settings.cdpAddress = text
                }
            }
        }

        RowLayout {
            width: parent.width
            anchors.bottom: parent.bottom

            Button {
                id: cancelBtn
                Layout.alignment: Qt.AlignLeft
                Material.background: Material.color(Material.Grey, Material.Shade300)
                Material.foreground: Material.color(Material.Grey, Material.Shade900)
                text: "Cancel"
                onClicked: dialog.reject()
            }

            Button {
                id: saveBtn
                Layout.alignment: Qt.AlignRight
                Material.background: Material.Indigo
                Material.foreground: Material.color(Material.Grey, Material.Shade50)
                text: "Save"
                onClicked: dialog.accept()
            }
        }

        onAccepted: {
            saveDate(noop)
        }

        onRejected: {
            loadData(noop)
        }
    }
}
