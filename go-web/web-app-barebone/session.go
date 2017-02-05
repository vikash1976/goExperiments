package main

import(
    "github.com/satori/go.uuid"
	"net/http"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
    c, err := req.Cookie("session")
    if err != nil {
        sID := uuid.NewV4()
        c = &http.Cookie {
            Name: "session",
            Value: sID.String(),
        }
        http.SetCookie(w, c)
    }
    var u user
    if un, ok := dbSession[c.Value]; ok {
        u = dbUsers[un]
    }
    
    return u
}

func alreadyLoggedIn(req *http.Request) bool {
    c, err := req.Cookie("session")
    if err != nil {
        return false
    }
    un := dbSession[c.Value]
    _, ok := dbUsers[un]
    return ok
}