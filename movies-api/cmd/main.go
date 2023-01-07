package main

import (
	"github.com/gorilla/mux"
	"log"
	"movies-api/pkg/routes"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRouter(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
