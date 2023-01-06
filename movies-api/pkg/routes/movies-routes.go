package routes

import (
	"github.com/gorilla/mux"
	"movies-api/pkg/controllers"
)

func Routes(router *mux.Router) {
	router.HandleFunc("/movies/", controllers.getMovies).Methods("GET")
	router.HandleFunc("/movies/", controllers.createMovies).Methods("POST")
	router.HandleFunc("/movies/{id}", controllers.getMovie).Methods("GET")
	router.HandleFunc("/movies/{id}", controllers.updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", controllers.deleteMovie).Methods("DELETE")
}
