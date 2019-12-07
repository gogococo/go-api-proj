package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Hello World!")
	})
	// func listen_and_serve, pass in address ex port 8080, and router (mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

// github.com/gorilla/mux
