package main

import (
	"github.com/gorilla/mux"
	"strings"
	"gopkg.in/yaml.v2"
	"fmt"
	"net/http"
)

/* If the URL is "http://localhost:8080", takes baseData
	If the URL is "http://localhost:8080/{key}", check the key.
		If the key is a type-name or move name, lists pokemons by name.
		If the key is a pokemon-name, gets pokemon by name.
 */
func GetByNameOrListByTypeOrMove(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := strings.ToLower(vars["key"])
	var str string

	if key == "" { //if there is not a key(http://localhost:8080), takes baseData
		b, err := yaml.Marshal(baseData)
		Check(err)
		str = string(b)
		fmt.Fprint(w, str)
	} else { //if there is a key
		var pokemons []Pokemon

		pokemons = FilterByTypeName(pokemonBD, key) //Filter pokemons by key(type-name). If key is type-name, returns function and does not continue.
		if len(pokemons) != 0 {
			WriteSimplePokemons(w, NewSimplePokemons(pokemons))
			return
		}
		pokemons = FilterByMoveName(pokemonBD, key) //Filter pokemons by key(move-name). If key is move-name, returns function and does not continue.
		if len(pokemons) != 0 {
			WriteSimplePokemons(w, NewSimplePokemons(pokemons))
			return
		}

		for _, x := range pokemonBD { //if the key is pokemon-name, lists the pokemon whose name is equal the key.
			if key == strings.ToLower(x.Name) {
				WriteSimplePokemon(w, NewSimplePokemon(x))
				return
			}
		}
	}
}
