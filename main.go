package main

import (
    "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome"))
}


func main() {
    http.HandleFunc("/home", homeHandler)
    http.ListenAndServe(":8080", nil)
}
