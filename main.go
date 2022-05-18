package main

// we will not be using a database for this, we will be using structs and slices to
import (
	"encoding/json" // encode data into json when sent into postman
	"fmt"           // to print stuff out to screen when conenctedf to server
	"log"           // log any errors if any errors are found connecting to server
	"math/rand"     // if user adds a new movie to the server you will need to create an id for it and this will help
	"net/http"      // allows us to create a server in golang
	"strconv"       // the id you will create will be an integer and this will help convert it into a string

	"github.com/gorilla/mux" // this is the external library we just installed we need to import it to use it.
)

type Movie struct { // remember to capitalize your first letter in the type of struct and the type of value in it for json to read it
	ID       string    `json:"id"` // will be used to encode into json by marhsalling
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"` //* is a pointer, pointing at Director struct
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie // slice of type movie has a zero/nil value

func getMovies(w http.ResponseWriter, r *http.Request) {
	// you are setting the content to allow it to convert it to json
	w.Header().Set("Content-Type", "application/json")
	// this encodes the data from the struct into json from the response you get
	// by passing w from reponse and the data in the movies variable that is available globally
	json.NewEncoder(w).Encode(movies) // this will return all movies
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// the id you are getting will be passed as a param that is inside mux.Vars()
	params := mux.Vars(r)
	// 	ranging over movies to search for the ID that is assigned to params if its there it will be deleted by appending
	for index, item := range movies {
		// if the ID you passed to params matches the id in the data to passed into []movies
		if item.ID == params["id"] {
			// you will delete that movie by id by appending the rest the movies not matching that id back into the []movie and replacing that one ID you matched
			movies = append(movies[:index], movies[index+1:]...) // the movies[index+1:]...) is attaching all the movies to movies but using the ... wihtout having to name them
			break                                                // this breaks out of the function once it has been completed.
		}
	}
	json.NewEncoder(w).Encode(movies) // once movie is deleted it will return the remaining movies
}

// remember you need to send and receive information via response and request.
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // this is getting the info from r *http.Request
	// you are ranging through the movies by ID and not the index so you throw it away but put a _ in its place since go doesnt allow it without it
	for _, item := range movies {
		if item.ID == params["id"] {
			// if the id in movies matches the id in params it will return that id through json
			json.NewEncoder(w).Encode(item) // this will return that one specific movie.
			return                          // to get out of function once condition is met
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	// always needed
	w.Header().Set("Content-Type", "application/json")
	// this var is only available within this function making a variable of movie of TYPE Movie
	var movie Movie
	// making a blank identifier and decode the body inside that movie from json to code golang can understand by giving it the address &of movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// creating a new ID using math rand pkg with an ID of 1 - 100000 that will then be converted into a string using strconv.Itoa
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	// now you are adding the created movie by appending the new var movie in this function to var movies []Movie on the global scope
	movies = append(movies, movie)
	// now you will convert it from go to json by encoding
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)

	//loop over the movies, range
	for index, item := range movies {
		//delete the mvoie with the ID that youve created
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			//add a new movie - the movie that we send in the body of postman
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func main() {
	r := mux.NewRouter() //is a function insde the mux gorilla library. r is not our new router

	// here we are adding to var movies by appending to it the data in the structs. the & is the address being passed onto the *Director pointer
	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movie Two", Director: &Director{Firstname: "Steven", Lastname: "Smith"}})
	// you will have 5 routes and 5 functions for each route yoyu will need to create them below look at your chart
	// gets all the movies by calling getMovies from the /movies route
	r.HandleFunc("/movies", getMovies).Methods("GET")
	// gets all movies by id calling getMovie singular from the movies/id route
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	// creates a movie by calling createMovie on the movies route
	r.HandleFunc("/movies", createMovie).Methods("POST")
	// updates a movie by calling updateMovie on the movies route
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	// deletes a movie by calling delteMovie on the movies route
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	// this starts the server first you need to print to screen if connection successful
	log.Fatal(http.ListenAndServe(":8000", r)) // you do need to put the : in from of 8000 because of localhost:8000 when you go to that address

}
