package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/html/home.html")
    })

    http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/html/about.html")
    })

    http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/html/contact.html")
    })

    http.HandleFunc("/stylesheets/main.css", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/css/main.css")
    })

    http.HandleFunc("/scripts/main.js", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "./static/js/main.js")
    })

    http.ListenAndServe(":8080", nil)
}
