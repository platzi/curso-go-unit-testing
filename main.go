package main

import (
	"catching-pokemons/controller"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/pokemon/{id}", controller.GetPokemon).Methods("GET")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Print("Error found")
	}
}
