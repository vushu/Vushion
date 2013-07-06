package main

import (
    "net/http"
    "github.com/gorilla/mux"
)
func init(){
    r := mux.NewRouter()
    r.HandleFunc("/", homeHandler)
    r.HandleFunc("/about", aboutHandler)
    r.HandleFunc("/games", gamesHandler)
    r.HandleFunc("/pixelart",pixelartHandler)
    r.PathPrefix("/static/").Handler(http.FileServer(http.Dir(".")))
    http.Handle("/",r)
}


