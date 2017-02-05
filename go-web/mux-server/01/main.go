package main

import (
    "net/http"
    "io"
    "html/template"
   
)

var tpl *template.Template
func init() {
    tpl = template.Must(template.ParseFiles("me.gohtml"))
    
}

func index(w http.ResponseWriter, r *http.Request){
    io.WriteString(w, r.Method)
}

func myself(w http.ResponseWriter, r *http.Request){
    //io.WriteString(w, "<html><body><h1>I am Vikash</h1></body></html>")
    tpl.ExecuteTemplate(w, "me.gohtml", "Vikash")
}

func aboutDog(w http.ResponseWriter, r *http.Request){
    io.WriteString(w, "This is my dog, Puffy")
}
func main () {
    http.HandleFunc("/", index)
    //http.HandleFunc("/me", myself)
    http.HandleFunc("/dog", aboutDog)
    http.Handle("/me", http.HandlerFunc(myself))
    
    http.ListenAndServe(":8080", nil)
}