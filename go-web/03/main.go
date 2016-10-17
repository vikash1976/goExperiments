package main

import (
    "net/http"

    "github.com/russross/blackfriday"
)
//run it from any folder to server its content
//go build main.go
//locate main.exe here from any other folder ./<<relative path to this>>/main.exe


func main() {
    http.HandleFunc("/markdown", GenerateMarkdown)
    http.Handle("/", http.FileServer(http.Dir("public"))) // Handle on / is catch all hence needs to be last handler
    http.ListenAndServe(":8080", nil)
}
//GenerateMarkdown handler for markdown route
func GenerateMarkdown(rw http.ResponseWriter, r *http.Request) {
    markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
    rw.Write(markdown)
}