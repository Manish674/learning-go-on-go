package routes

import (
	"github.com/gorilla/mux"
	"movies-api/pkg/controllers"
)

func RegisterRouter(router *mux.Router) {
	router.HandleFunc("/movies/", movie_controller.getMovies).Methods("GET")
	router.HandleFunc("/movies/", movie_controller.createMovies).Methods("POST")
	router.HandleFunc("/movies/{id}", movie_controller.getMovie).Methods("GET")
	router.HandleFunc("/movies/{id}", movie_controller.updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", movie_controller.deleteMovie).Methods("DELETE")
}
