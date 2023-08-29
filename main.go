package main


import (
	"fmt"
	"net/http"
	"encoding/json"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
	"log"
)


// TYPE DELCARATIONS
type Movie struct {

	Director *Director `json:"director"`
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	
}


type Director struct {

	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`


}


var movies []Movie


// FUNCTION DECLARATIONS


func getMovies(w http.ResponseWriter, r *http.Request) {


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)




}


func deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}


func getMovie(w http.ResponseWriter, r *http.Request) {


	w.Header().Set("Content-Type", "application/json")


	params := mux.Vars(r)
	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}


func createMovie (w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}


func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for i, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}


	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = params["id"]
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func main() {



	r := mux.NewRouter()


	movies = append(movies, Movie{
		Director: &Director{
			Firstname: "Sean",
			Lastname: "Bain",
		},
		ID: "1",
		Isbn: "4382777",
		Title: "Movie1",
	})

	movies = append(movies, Movie{
		Director: &Director{
			Firstname: "John",
			Lastname: "Smith",
		},
		ID: "2",
		Isbn: "432123",
		Title: "Movie2",
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods(("GET"))
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")



	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", r))

}