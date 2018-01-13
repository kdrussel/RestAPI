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
	ID 		string `json:"id"`
	Name 	string `json:"name"`
	Alias 	string `json:"alias"`
	Birthday string `json:"birthday"`

}

// Init books var as a slice Hero struct
var heroes []Hero

// Get all heroes
func getHeroes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(heroes)
}

// Get single hero
func getHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find with id
	for _, item := range heroes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Hero{})
}

// Create a new hero
func createHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var hero Hero
	_ = json.NewDecoder(r.Body).Decode(&hero)
	hero.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID
	heroes = append(heroes, hero)
	json.NewEncoder(w).Encode(hero)
}

// Update hero
func updateHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range heroes {
		if item.ID == params["id"] {
		heroes = append(heroes[:index], heroes[index+1:]...)
		break
		}
	}
	json.NewEncoder(w).Encode(heroes)
}

// Delete hero
func deleteHero(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range heroes {
		if item.ID == params["id"] {
		heroes = append(heroes[:index], heroes[index+1:]...)
		var hero Hero
		_ = json.NewDecoder(r.Body).Decode(&hero)
		hero.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID
		heroes = append(heroes, hero)
		json.NewEncoder(w).Encode(hero)
		return
		}
	}
	json.NewEncoder(w).Encode(heroes)
}

func main() {
	// Init router
	r := mux.NewRouter()

	// Mock data - add DB
	heroes = append(heroes, Hero{ID: "1", Name:"Bruce Wayne", Alias: "Batman", Birthday: "Unknown"})

	// Route handlers / Endpoints
	r.HandleFunc("/api/superheroes", getHeroes).Methods("GET")
	r.HandleFunc("/api/superheroes/{id}", getHero).Methods("GET")
	r.HandleFunc("/api/superheroes", createHero).Methods("POST")
	r.HandleFunc("/api/superheroes/{id}", updateHero).Methods("PUT")
	r.HandleFunc("/api/superheroes/{id}", deleteHero).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", r))
}