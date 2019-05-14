import QtQuick 2.4

ListModel {
    ListElement {
         name: "muesli"
         messageid: "901223685058703361"
         author: "mueslix"
         actorname: "eve"
         avatar: "https://pbs.twimg.com/profile_images/779041781413507072/TaqJsdzS_normal.jpg"
         body: "This is a very, very long test post, that should probably get word-wrapped. But does it work?"
         createdat: "now"
         reply: false
         mention: false
         like: true
         followed: false
    }
    ListElement {
         name: "somebody"
         messageid: "901223685058703361"
         author: "somebody"
         actorname: "fribbledom"
         avatar: "https://pbs.twimg.com/profile_images/779041781413507072/TaqJsdzS_normal.jpg"
         body: "Please boost it, it's my birthday"
         createdat: "now"
         reply: false
         forward: true
         mention: false
         followed: false
    }
}
