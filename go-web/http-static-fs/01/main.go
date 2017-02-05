package main

import (
    "net/http"
    "log"
    "html/template"
    "io"
)

func main() {
    http.HandleFunc("/", foo)
    http.HandleFunc("/dog/", dog)
    http.HandleFunc("/dog.jpg", dogPic)
    log.Fatal(http.ListenAndServe(":8080", nil))
    
}

func foo(w http.ResponseWriter, req *http.Request) {
    io.WriteString(w, "Ran Foo")
    //http.ServeFile(w, req, ".")
}

func dog(w http.ResponseWriter, req *http.Request) {
    log.Println("In dog function")
    tpl, err := template.ParseFiles("dog.gohtml")
    if err != nil {
        log.Fatalln(err)
    }
    err = tpl.ExecuteTemplate(w, "dog.gohtml", nil)
    if err != nil {
        log.Fatalln(err)
    }
}

func dogPic(w http.ResponseWriter, req *http.Request) {
    http.ServeFile(w, req, "dog.jpg")
}