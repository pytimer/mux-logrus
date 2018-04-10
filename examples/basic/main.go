package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pytimer/mux-logrus"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello!"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", index).Methods(http.MethodGet)

	r.Use(muxlogrus.NewLogger().Middleware)

	address := ":8990"
	log.Printf("Listen on address %s", address)
	http.ListenAndServe(address, r)
}
