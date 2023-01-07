package controllers

import (
	// "movies-api/pkg/models"
	"encoding/json"
	"fmt"
	"movies-api/pkg/models"
	"net/http"

	"github.com/gorilla/mux"
)

// "movies-api/pkg/models"
// "movies-api/pkg/config"

// mongodb helper method for movie operations
// func insertMovie(movie model.Movie) {
// }

func GetActorById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result := *models.GetActor(params["id"])
	fmt.Print("\n result --> ", result)
	res, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// func CreateMovie(w *http.ResponseWriter, r *http.Request) {
// }

// func GetMovie(w *http.ResponseWriter, r *http.Request) {
// }

// func DeleteMovie(w *http.ResponseWriter, r *http.Request) {
// }

// func UpdateMovie(w *http.ResponseWriter, r *http.Request) {
// }
