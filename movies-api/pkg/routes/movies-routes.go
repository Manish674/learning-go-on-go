package routes

import (
	"github.com/gorilla/mux"
	"movies-api/pkg/controllers"
)

func RegisterRouter(router *mux.Router) {
	router.HandleFunc("/movies/actor/{id}", controllers.GetActorById).Methods("GET")
	// router.HandleFunc("/movies/", movie_controller.CreateMovie).Methods("POST")
	// router.HandleFunc("/movies/{id}", movie_controller.GetMovie).Methods("GET")
	// router.HandleFunc("/movies/{id}", movie_controller.UpdateMovie).Methods("PUT")
	// router.HandleFunc("/movies/{id}", movie_controller.DeleteMovie).Methods("DELETE")
}
