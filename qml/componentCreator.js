function createMessagePopup(parent, model) {
    var component = Qt.createComponent("MessagePopup.qml")
    var popup = component.createObject(parent, {
        "message": model
    })

    if (popup == null) {
        console.log("Error creating MessagePopup")
    }
    return popup
}

function createSharePopup(parent, model) {
    var component = Qt.createComponent("SharePopup.qml")
    var popup = component.createObject(parent, {
        "message": model
    })

    if (popup == null) {
        console.log("Error creating SharePopup")
    }
    return popup
}

function createDeletePopup(parent, model) {
    var component = Qt.createComponent("DeletePopup.qml")
    var popup = component.createObject(parent, {
        "message": model
    })

    if (popup == null) {
        console.log("Error creating DeletePopup")
    }
    return popup
}

function createConversationPopup(parent, model) {
    var component = Qt.createComponent("ConversationPopup.qml")
    var popup = component.createObject(parent, {
        "message": model
    })

    if (popup == null) {
        console.log("Error creating ConversationPopup")
    }
    return popup
}

function createAccountPopup(parent) {
    var component = Qt.createComponent("AccountPopup.qml")
    var popup = component.createObject(parent, {})

    if (popup == null) {
        console.log("Error creating AccountPopup")
    }
    return popup
}
