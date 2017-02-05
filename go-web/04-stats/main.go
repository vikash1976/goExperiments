package main

import (
    "net/http"
    "github.com/codegangsta/negroni"
    "github.com/thoas/stats"
    "encoding/json"
)

func main() {
    middleware := stats.New()

    mux := http.NewServeMux()

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte("{\"hello\": \"world\"}"))
    })

    mux.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        stats := middleware.Data()

        b, _ := json.Marshal(stats)

        w.Write(b)
    })

    n := negroni.Classic()
    n.Use(middleware)
    n.UseHandler(mux)
    n.Run(":3000")
}