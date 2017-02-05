package main

import (
    "net/http"
    "log"
    "html/template"
    
)

var tpl *template.Template
func init() {
    tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}
func main() {
    
    fs := http.StripPrefix("/public", http.FileServer(http.Dir("public")))
    
    http.Handle("/public/", fs)
    http.HandleFunc("/", dog)
    log.Fatal(http.ListenAndServe(":8080", nil))
    
}

func dog(w http.ResponseWriter, req *http.Request) {
    log.Println("In dog")
    err := tpl.Execute(w, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    
}