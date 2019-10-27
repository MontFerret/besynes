import QtQuick 2.13
import QtQuick.Controls 2.13
import QtQuick.Controls.Material 2.13
import QtQuick.Layouts 1.12

Dialog {
    id: root
    modal: true
    standardButtons: Dialog.Cancel | Dialog.Save
    anchors.centerIn: parent
}
