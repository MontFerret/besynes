import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13
import QtQuick.Layouts 1.12

Item {
    id: root
    width: 400
    height: 200
    anchors.centerIn: parent

    readonly property var open: (msg) => {
        if (!msg) {
            return;
        }

        if (msg.title) {
            dialog.title = msg.title;
        }

        let text = "";

        if (typeof msg.body === "string") {
            text = msg.body;
        } else if (msg.body instanceof Error) {
            text = msg.body.message;
        } else if (typeof msg.body === "object") {
            text = JSON.stringify(msg.body);
        } else if (msg.body) {
            text = msg.body.toString();
        }


        dialogText.text = text;

//        const type = (msg.type || "info").toLowerCase();

//        switch (type) {
//            case "error": {

//                break;
//            }
//            case "warn": {
//                break;
//            }
//            case "info": {
//                break;
//            }
//            default: {
//                break;
//            }
//        }

        dialog.open()
    }

    Dialog {
        id: dialog
        width: root.width
        height: root.height
        visible: false

        RowLayout {
            width: parent.width

            Text {
                id: dialogText
                text: ""
            }
        }

        RowLayout {
            width: parent.width
            anchors.bottom: parent.bottom

            Button {
                id: okBtn
                Layout.alignment: Qt.AlignCenter
                Material.background: Material.Indigo
                Material.foreground: Material.color(Material.Grey, Material.Shade50)
                text: "Ok"
                onClicked: dialog.accept()
            }
        }
    }
}
