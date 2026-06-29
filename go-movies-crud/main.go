package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/google/uuid"
	
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


func deleteMovie (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	params := mux.Vars(r)
	
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	if err := json.NewEncoder(w).Encode(movies); err != nil {
		fmt.Println(err);
	}

}


func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _,item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
		
	}

	http.Error(w, "Movie not found", http.StatusNotFound)

}

func createMovie (w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	movie.ID = uuid.New().String()
	movies = append(movies, movie)
	
	if err := json.NewEncoder(w).Encode(movie); err != nil {
		fmt.Println(err);
	}

	


} 


func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

    params := mux.Vars(r)

	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index + 1:]...)
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return

		}
	}



}




var movies []Movie

func main(){
	r := mux.NewRouter()
    movies = append(movies, Movie{ID:"1", Title:"Movie One", Director:&Director{Firstname:"Jhon" , Lastname :"Doe"}})
	movies = append(movies, Movie{ID:"2", Title:"Movie Two", Director:&Director{Firstname:"Steve" , Lastname :"Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET");
    r.HandleFunc("/movies/{id}",deleteMovie).Methods("DELETE");
    r.HandleFunc("/movies/{id}", getMovie).Methods("GET");
	r.HandleFunc("/movies",createMovie).Methods("POST");
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT");


	fmt.Println("Starting server at port 8000")

	if err := http.ListenAndServe(":8000", r); err != nil {
		log.Fatal(err)
	}
	

}