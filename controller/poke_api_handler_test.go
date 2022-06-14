package controller

import (
	"catching-pokemons/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/require"
)

func TestGetPokemonFromPokeApiSuccess(t *testing.T) {
	c := require.New(t)

	pokemon, err := GetPokemonFromPokeApi("bulbasaur")
	c.NoError(err)

	body, err := ioutil.ReadFile("samples/poke_api_read.json")
	c.NoError(err)

	var expected models.PokeApiPokemonResponse

	err = json.Unmarshal([]byte(body), &expected)
	c.NoError(err)

	c.Equal(expected, pokemon)
}

func TestGetPokemonFromPokeApiSuccessWithMocks(t *testing.T) {
	c := require.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "balbasaur"

	body, err := ioutil.ReadFile("samples/poke_api_response.json")
	c.NoError(err)

	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	httpmock.RegisterResponder("GET", request, httpmock.NewStringResponder(200, string(body)))

	pokemon, err := GetPokemonFromPokeApi(id)
	c.NoError(err)

	var expected models.PokeApiPokemonResponse

	err = json.Unmarshal([]byte(body), &expected)
	c.NoError(err)

	c.Equal(expected, pokemon)
}

func TestGetPokemonFromPokeApiInternalServerError(t *testing.T) {
	c := require.New(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	id := "balbasaur"

	body, err := ioutil.ReadFile("samples/poke_api_response.json")
	c.NoError(err)

	request := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", id)

	httpmock.RegisterResponder("GET", request, httpmock.NewStringResponder(500, string(body)))

	_, err = GetPokemonFromPokeApi(id)
	c.NotNil(err)
	c.EqualError(ErrPokeApiFailure, err.Error())
}

func TestGetPokemon(t *testing.T) {
	c := require.New(t)

	r, err := http.NewRequest("GET", "/pokemon/{id}", nil)
	c.NoError(err)

	w := httptest.NewRecorder()

	vars := map[string]string{
		"id": "bulbasaur",
	}

	r = mux.SetURLVars(r, vars)

	GetPokemon(w, r)

	expectedBodyResponse, err := ioutil.ReadFile("samples/api_response.json")
	c.NoError(err)

	var expectedPokemon models.Pokemon

	err = json.Unmarshal([]byte(expectedBodyResponse), &expectedPokemon)
	c.NoError(err)

	var actualPokemon models.Pokemon

	err = json.Unmarshal([]byte(w.Body.Bytes()), &actualPokemon)
	c.NoError(err)

	c.Equal(http.StatusOK, w.Code)
	c.Equal(expectedPokemon, actualPokemon)
}
