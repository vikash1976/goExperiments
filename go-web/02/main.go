package main

import (
    "fmt"
    "net/http"
     "github.com/vikash1976/goExperiments/go-web/02/page"
     "regexp"     
)

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
//function with closure to bring title check and error handling at one place
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        m := validPath.FindStringSubmatch(r.URL.Path)
        if m == nil {
            http.NotFound(w, r)
            return
        }
        fn(w, r, m[2])
    }
}
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
    fmt.Fprintf(w, "I love %s!", r.URL.Path[1:])
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
   
    p, err:= page.LoadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/"+title, http.StatusFound)
        return
    }
    page.RenderTemplatePage(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
   
    p, err := page.LoadPage(title)
    if err != nil {
        p = &page.Page{Title: title}
    }
    page.RenderTemplatePage(w, "edit", p)
    
}
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    //without this check save was accessible as GET, terminating the page content to 0 byte.
    if r.Method == "GET" {
        http.Error(w, "Invalid HTTP Method on Save", http.StatusInternalServerError)
        return
    }
    
    body := r.FormValue("body")
    p := &page.Page{Title: title, Body: []byte(body)}
    err := p.Save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main(){
    http.HandleFunc("/view/",  makeHandler(viewHandler))
    http.HandleFunc("/edit/",  makeHandler(editHandler))
    http.HandleFunc("/save/",  makeHandler(saveHandler))
    http.ListenAndServe(":8070", nil)
}