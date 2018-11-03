package sessions

import "github.com/gorilla/sessions"

var Store = sessions.NewCookieStore([]byte("something-secretive-is-what-a-gorrilla-needs"))
