package main

import (
	"net/http"
	"os"

	"github.com/gophergala2016/chopher/api"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	r := mux.NewRouter()
	r.StrictSlash(true)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/"))).Methods("GET")
	r.HandleFunc("/upload", api.FileUploadHandler).Methods("POST")
	http.ListenAndServe(":"+port, handlers.LoggingHandler(os.Stdout, r))
}
