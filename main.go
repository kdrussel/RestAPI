package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

// Heroes Struct (Model)
type Hero struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Alias string `json:"alias"`
	Birthday string `json:"birthday"`

}

func main() {
	// Init router
	r := mux.NewRouter()

	// Route handlers / Endpoints
	r.HandleFunc("/api/superheroes", getHeroes).Methods("GET")
	r.HandleFunc("/api/superheroes/{id}", getHero).Methods("GET")
	r.HandleFunc("/api/superheroes", createHero).Methods("POST")
	r.HandleFunc("/api/superheroes/{id}", updateHero).Methods("PUT")
	r.HandleFunc("/api/superheroes/{id}", deleteHero).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000". r))
}