function loadComponent(qml) {
    var component = Qt.createComponent(qml)
    if (component.status != Component.Ready &&
        component.status == Component.Error) {
        console.debug("Error loading component (" + qml + "): " + component.errorString())
        return null
    }

    return component;
}

function createMessagePopup(parent, model) {
    var component = loadComponent("MessagePopup.qml")
    if (component == null) {
        return null
    }

    accountBridge.attachments.clear()

    var popup = component.createObject(parent, {
        "message": model
    })
    if (popup == null) {
        console.log("Error creating MessagePopup")
    }
    return popup
}

function createSharePopup(parent, model) {
    var component = loadComponent("SharePopup.qml")
    if (component == null) {
        return null
    }

    var popup = component.createObject(parent, {
        "message": model
    })
    if (popup == null) {
        console.log("Error creating SharePopup")
    }
    return popup
}

function createDeletePopup(parent, model) {
    var component = loadComponent("DeletePopup.qml")
    if (component == null) {
        return null
    }

    var popup = component.createObject(parent, {
        "message": model
    })
    if (popup == null) {
        console.log("Error creating DeletePopup")
    }
    return popup
}

function createConversationPopup(parent, model) {
    var component = loadComponent("ConversationPopup.qml")
    if (component == null) {
        return null
    }

    var popup = component.createObject(parent, {
        "message": model
    })
    if (popup == null) {
        console.log("Error creating ConversationPopup")
    }
    return popup
}

function createAccountPopup(parent) {
    var component = loadComponent("AccountPopup.qml")
    if (component == null) {
        return null
    }

    var popup = component.createObject(parent, {})
    if (popup == null) {
        console.log("Error creating AccountPopup")
    }
    return popup
}

function createMediaPopup(parent, model) {
    var component = loadComponent("MediaPopup.qml")
    if (component == null) {
        return null
    }

    var popup = component.createObject(parent, {
        "url": model
    })
    if (popup == null) {
        console.log("Error creating MediaPopup")
    }
    return popup
}
