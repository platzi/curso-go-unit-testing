package controller

import (
	"catching-pokemons/models"
	"catching-pokemons/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	response, err := http.Get(request)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var apiPokemon models.PokeApiPokemonResponse

	err = json.Unmarshal(body, &apiPokemon)
	if err != nil {
		log.Fatal(err)
	}

	parsedPokemon, err := util.ParsePokemon(apiPokemon)
	if err != nil {
		respondwithJSON(w, http.StatusInternalServerError, fmt.Sprintf("error found: %s", err.Error()))
	}

	respondwithJSON(w, http.StatusOK, parsedPokemon)
}
