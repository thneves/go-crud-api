package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	route := mux.NewRouter()

	movies = append(movies, Movie{
		ID:       "1",
		Isbn:     "4800",
		Title:    "Now You See Me",
		Director: &Director{"John", "Doe"},
	})

	movies = append(movies, Movie{
		ID:       "2",
		Isbn:     "4801",
		Title:    "Dune",
		Director: &Director{"Denis", "Villeneuve"},
	})

	route.HandleFunc("/movies", getMovies).Methods("GET")
	route.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	route.HandleFunc("/movies", createMovie).Methods("POST")
	route.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	route.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server started at port 8000")
	log.Fatal(http.ListenAndServe(":8000", route))
}

func getMovies(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(write)
	encoder.Encode(movies)
}

func deleteMovie(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	encoder := json.NewEncoder(write)
	encoder.Encode(movies)
}

func getMovie(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range movies {
		if item.ID == params["id"] {
			encoder := json.NewEncoder(write)
			encoder.Encode(item)
			return
		}
	}
}

func createMovie(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(request.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn((100000)))
	movies = append(movies, movie)
	encoder := json.NewEncoder(write)
	encoder.Encode(movie)
}

func updateMovie(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			_ = json.NewDecoder(request.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			encoder := json.NewEncoder(write)
			encoder.Encode(movies)
			return
		}
	}
}
