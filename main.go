package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"math/rand"

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

// Get all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")
	// Return the movies
	json.NewEncoder(w).Encode(movies)
}

// Delete a movie
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")
	// Get the params from the request
	params := mux.Vars(r)
	// Loop through the movies
	for index, item := range movies {
		// If the movie ID is equal to the params ID
		if item.ID == params["id"] {
			// Remove the movie from the movies slice
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	// Return the movies
	json.NewEncoder(w).Encode(movies)
}

// Get a single movie
func getMovie(w http.ResponseWriter, r *http.Request) {
	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")
	// Get the params from the request
	params := mux.Vars(r)
	// Loop through the movies
	for _, item := range movies {
		// If the movie ID is equal to the params ID
		if item.ID == params["id"] {
			// Return the movie
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Create a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Create a new movie
	var movie Movie
	// Decode the request body and assign it to the movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// Set the ID to a random number
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	// Append the movie to the movies slice
	movies = append(movies, movie)
	// Return the new movie
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	// Set the content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Get the params from the request
	params := mux.Vars(r)

	// Loop through the movies
	for index, item := range movies {
		// If the movie ID is equal to the params ID
		if item.ID == params["id"] {
			// Remove the movie from the movies slice
			movies = append(movies[:index], movies[index+1:]...)
			// Create a new movie
			var movie Movie
			// Decode the request body and assign it to the movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			// Set the ID to the params ID
			movie.ID = params["id"]
			// Append the movie to the movies slice
			movies = append(movies, movie)
			// Return the new movie
			json.NewEncoder(w).Encode(movie)
			// Return the movies
			return
		}
	}
}

func main() {
	// Init the router
	r := mux.NewRouter()

	// Mock data
	movies = getMockDatas()

	// Route handlers / Endpoints
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// Start the server
	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Mock data
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
