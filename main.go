package main 
// we will not be using a database for this, we will be using structs and slices to 
import (
	"fmt" // to print stuff out to screen when conenctedf to server
	"log" // log any errors if any errors are found connecting to server
	"encoding/json" // encode data into json when sent into postman
	"math/rand" // if user adds a new movie to the server you will need to create an id for it and this will help
	"net/http" // allows us to create a server in golang
	"strconv" // the id you will create will be an integer and this will help convert it into a string
	"github.com/gorilla/mux" // this is the external library we just installed we need to import it to use it.
)

type Movie struct { // remember to capitalize your first letter in the type of struct and the type of value in it for json to read it
	ID string `json:"id"` // will be used to encode into json by marhsalling
	Isbn string `json: "isbn"`
	Title string `json: "title"`
	Director *Director `json:"director"` //* is a pointer, pointing at Director struct
} 

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie


func main(){
	r := mux.NewRouter() //is a function insde the mux gorilla library. r is not our new router

	// you will have 5 routes and 5 functions for each route yoyu will need to create them below look at your chart
		// gets all the movies by calling getMovies from the /movies route
	r.HandleFunc("/movies", getMovies).Methods("GET")
		// gets all movies by id calling getMovie singular from the movies/id route
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
		// creates a movie by calling createMovie on the movies route
	r.HandleFunc("/movies", createMovie).Methods("POST")
		// updates a movie by calling updateMovie on the movies route
	r.HandleFunc("/movies", updateMovie).Methods("PUT")
		// deletes a movie by calling delteMovie on the movies route
	r.HandleFunc("/movies", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	// this starts the server first you need to print to screen if connection successful 
	log.Fatal(http.ListenAndServe(":8000", r)) // you do need to put the : in from of 8000 because of localhost:8000 when you go to that address

	
}