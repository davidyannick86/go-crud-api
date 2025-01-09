package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Year     string    `json:"year"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.NewRouter()

	movies = getMockDatas()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getMockDatas() []Movie {
	movies = append([]Movie{}, Movie{
		ID:       "1",
		Isbn:     "438227",
		Title:    "The Godfather",
		Year:     "1972",
		Director: &Director{Firstname: "Francis Ford", Lastname: "Coppola"},
	})

	movies = append(movies, Movie{
		ID:       "2",
		Isbn:     "454545",
		Title:    "Goodfellas",
		Year:     "1990",
		Director: &Director{Firstname: "Martin", Lastname: "Scorsese"},
	})

	movies = append(movies, Movie{
		ID:       "3",
		Isbn:     "454545",
		Title:    "The Shawshank Redemption",
		Year:     "1994",
		Director: &Director{Firstname: "Frank", Lastname: "Darabont"},
	})
	return movies
}
