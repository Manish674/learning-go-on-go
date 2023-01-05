package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	ID          string    `json:"id"`
	MovieNumber string    `json:"movieNumber"`
	Title       string    `json:"title"`
	Director    *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

var movies []Movie

// req  response
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	// index no.., item will be the single movie item
	for index, item := range movies {
		if item.ID == params["id"] {
			// basically deleteing the slice with index no
			// movies[index+1:]... is basically shifting all the movies to left side maintaining order
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies[index] = movie
			// item = movie
			json.NewEncoder(w).Encode(movies)
		}
	}
	// find the movie
}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", MovieNumber: "f23wa", Title: "Shamitabh", Director: &Director{FirstName: "R", LastName: "Balakrishnan", Age: 57}})
	movies = append(movies, Movie{ID: "2", MovieNumber: "f22df2", Title: "Fight Club", Director: &Director{FirstName: "David", LastName: "Fincher", Age: 57}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")

	fmt.Printf("Listenting server at localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
