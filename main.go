package main

import (
	"fmt"
	"log"
	"net/http"

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
	route.HandleFunc("/movies/{id}", getMovies).Methods("GET")
	route.HandleFunc("/movies", createMovie).Methods("POST")
	route.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	route.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Server started at port 8000")
	log.Fatal(http.ListenAndServe(":8000", route))
}
