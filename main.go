package main

import(
	"encoding/json"
	"log"
	"net/http"
	"math/rand"
	"strconv"
	"github.com/gorilla/mux"
)

func main() {
	// Init router
	r := mux.NewRouter()

	// Route handlers / Endpoints
	r.HandleFunc("/api/superheroes", getHeroes).Methods("GET")
	r.HandleFunc("/api/superheroes/{id}", getHero).Methods("GET")
	r.HandleFunc("/api/superheroes", createHero).Methods("POST")
	r.HandleFunc("/api/superheroes/{id}", updateHeroes).Methods("PUT")
	r.HandleFunc("/api/superheroes/{id}", deleteHeroes).Methods("DELETE")

}