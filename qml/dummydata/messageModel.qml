import QtQuick 2.4

ListModel {
    ListElement {
         name: "Christian Muehlhaeuser"
         messageid: "901223685058703361"
         author: "mueslix"
         avatar: "https://pbs.twimg.com/profile_images/779041781413507072/TaqJsdzS_normal.jpg"
         body: "This is a very, very long test post, that should probably get word-wrapped. But does it work?"
         createdat: "now"
         reply: false
         forward: false
         mention: false
         like: false
    }
    ListElement {
         avatar: "https://pbs.twimg.com/profile_images/779041781413507072/TaqJsdzS_normal.jpg"
         createdat: "now"
         actor: "afriend"
         actorname: "Alex Friend"
         reply: false
         forward: false
         mention: false
         like: false
         followed: true
    }
    ListElement {
         name: "Some Guy With A Really Really Long Name"
         messageid: "901223685058703361"
         author: "someguy"
         avatar: "https://pbs.twimg.com/profile_images/707382834827120640/R-Eb9YZB_normal.jpg"
         body: "This is a response"
         createdat: "now"
         reply: true
         replytoid: "901223685058703361"
         replytoauthor: "mueslix"
         forward: false
         mention: false
         like: false
         liked: true
    }
    ListElement {
         name: "Dummy User"
         messageid: "901223685058703361"
         author: "dummy"
         avatar: "https://pbs.twimg.com/profile_images/707382834827120640/R-Eb9YZB_normal.jpg"
         body: "This is a very short test post with a link: <a href=\"http://chris.de\">http://chris.de</a>"
         createdat: "now"
         actor: "mueslix"
         actorname: "Christian Muehlhaeuser"
         reply: false
         forward: true
         mention: false
         like: false
         shared: true
         mediapreview: "https://pbs.twimg.com/media/DIfdvcxXkAUXAvs.jpg"
    }
    ListElement {
         name: "Another User"
         messageid: "901223685058703361"
         author: "anotheruser"
         avatar: "https://pbs.twimg.com/profile_images/658723533845954560/noXJEv_a_normal.jpg"
         body: "Yet another test post. Lorem Ipsum Yada Yada @mueslix"
         createdat: "now"
         actor: "mueslix"
         actorname: "Christian Muehlhaeuser"
         reply: false
         forward: false
         mention: true
         like: false
    }
    ListElement {
         name: "This Poster"
         messageid: "901223685058703361"
         author: "posty"
         avatar: "https://pbs.twimg.com/profile_images/293948630/twitter_icon_normal.JPG"
         body: "I can't come up with any more mocking data now, really. This is enough."
         createdat: "now"
         actor: "mueslix"
         actorname: "Christian Muehlhaeuser"
         reply: false
         forward: false
         mention: false
         like: true
    }
}
