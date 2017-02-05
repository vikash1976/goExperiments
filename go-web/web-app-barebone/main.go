package main

import(
    "net/http"
    "log"
    "html/template"
    "github.com/satori/go.uuid"
    "golang.org/x/crypto/bcrypt"
    "fmt"
)

type user struct {
    UserName string
    Password []byte
    First string
}
/*type loginError struct {
    Error string
}
var loginErr loginError*/
var dbUsers = map [string]user{}
var dbSession = map[string]string{}
var tpl *template.Template

func init() {
    tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
    http.HandleFunc("/", handleRoot)
    http.HandleFunc("/signup", signup)
    http.HandleFunc("/login", login)
    http.HandleFunc("/logout", logout)
    http.HandleFunc("/context", showContext)
    http.Handle("/favicon.ico", http.NotFoundHandler())
    
    log.Fatalln(http.ListenAndServe(":8080", nil))
}

func handleRoot(w http.ResponseWriter, req *http.Request) {
    log.Println("In  handleRoot")
    u := getUser(w, req)
    handleError(tpl.ExecuteTemplate(w, "index.gohtml", u))
    
}

func showContext(w http.ResponseWriter, req *http.Request) {
    ctx := req.Context()
    log.Println(ctx)
    fmt.Fprintln(w, ctx)    
}

func signup(w http.ResponseWriter, req *http.Request) {
    /*if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}*/
    if(req.Method == http.MethodPost){
        un := req.FormValue("username")
        p := req.FormValue("password")
        f := req.FormValue("firstname")
    
        if _, ok := dbUsers[un]; ok {
            http.Error(w, "User already taken", http.StatusForbidden)
            return
        }
        /*sID := uuid.NewV4()
        c := &http.Cookie{
            Name: "session",
            Value: sID.String(),
        }
        http.SetCookie(w, c)
        dbSession[c.Value] = un*/
        pbs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
        if err != nil {
            http.Error(w, "Internal Sever Error", http.StatusInternalServerError)
            return
        }
        u := user {un, pbs, f }
        dbUsers[un] = u
        http.Redirect(w, req, "/login", http.StatusSeeOther)
        return
    }
    handleError(tpl.ExecuteTemplate(w, "signup.gohtml", nil))
}

func login(w http.ResponseWriter, req *http.Request) {
    if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
    
    if(req.Method == http.MethodPost){
        un := req.FormValue("username")
        p := req.FormValue("password")
        
    
       u, ok := dbUsers[un]
       if !ok {
            http.Error(w, `Username or password is incorrect <a href="/login">Login</a>`, http.StatusForbidden)
            //loginErr = loginError {"Username or password is incorrect"}
            //http.Redirect(w, req, "/login", http.StatusSeeOther)
            return
        }
        err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
        if err != nil {
            http.Error(w, `Username or password is incorrect <a href="/login">Login</a>`, http.StatusForbidden)
            //loginErr = loginError {"Username or password is incorrect"}
            //http.Redirect(w, req, "/login", http.StatusSeeOther)
            return
        }
        sID := uuid.NewV4()
        c := &http.Cookie{
            Name: "session",
            Value: sID.String(),
        }
        http.SetCookie(w, c)
        dbSession[c.Value] = un
        //loginErr = loginError {}
        http.Redirect(w, req, "/", http.StatusSeeOther)
        return
    }
    
    handleError(tpl.ExecuteTemplate(w, "login.gohtml", nil))
}
func logout(w http.ResponseWriter, req *http.Request) {
    if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
    c, _ := req.Cookie("session")
    delete(dbSession, c.Value)
    
    c = &http.Cookie {
        Name: "session",
        Value: "",
        MaxAge: -1,
    }
    
    http.SetCookie(w, c)
    
    http.Redirect(w, req, "/login", http.StatusSeeOther)
    
}

func handleError(err error){
    if err != nil {
        log.Fatalln(err)
    }
}


