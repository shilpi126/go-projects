package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type Movie struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Director *Director `jsonL"director"`
}

type Director struct{
	Firstname string `json: "firstname"`
	Lastname string `json: "lastname"`
}

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		http.Error(w, "Failed to encode movies", http.StatusInternalServerError)
		log.Printf("encode movies: %v", err)
		
	}
	
}




var movies []Movie

func main(){
	r := mux.NewRouter()
    movies = append(movies, Movie{ID:"1", Title:"Movie One", Director:&Director{Firstname:"Jhon" , Lastname :"Doe"}})
	movies = append(movies, Movie{ID:"2", Title:"Movie Two", Director:&Director{Firstname:"Steve" , Lastname :"Smith"}})
	r.HandleFunc("/movies", getMovies).Methods("GET");


	fmt.Println("Starting server at port 8000")

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
	

}