package main

import (
	"net/http"
    "fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func main() {
    fmt.Println("Server up at :8080")
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	http.ListenAndServe(":8080", router)
}
