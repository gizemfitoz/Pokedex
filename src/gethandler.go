package main

import (
	"net/http"
	"net/url"
	"strings"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"fmt"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	s := r.URL.String()
	u, err := url.Parse(s)
	Check(err)
	var obj interface{}
	if u.RawQuery != "" { //if URL contains query
		m, _ := url.ParseQuery(u.RawQuery) //map for query
		for k := range m {
			kk := strings.ToLower(string(k))
			if kk == "name" { //if the query key is "name", checks query value is which type of name(pokemon-name, type-name or move-name) and gets info according to that type of name.
				obj = GetByName(m[k][0])
			}
		}
	} else { //if URL does not contain query
		vars := mux.Vars(r)
		key := vars["key"]

		if key != "" { //if there is a key
			obj = GetByName(key) //checks the key is which type of name(pokemon-name, type-name or move-name) and gets info according to that type of name.
		}
	}

	ret, _ := yaml.Marshal(obj)
	fmt.Fprint(w, string(ret))
}

//Checks the parameter is which type of name(pokemon-name, type-name or move-name) and gets info by the type-name.
func GetByName(key string) interface{} {
	key = strings.ToLower(key)
	for _, x := range pokemonBD { //if the key is a pokemon-name, gets the that pokemon.
		if key == strings.ToLower(x.Name) {
			return GetByPokemonName(key, x)
		}
	}
	for _, typee := range typeBD { //if the key is a type-name, gets the that type.
		if key == strings.ToLower(typee.Name) {
			return typee
		}
	}
	for _, move := range moveBD { //if the key is a move-name, gets the that move.
		if key == strings.ToLower(move.Name) {
			return move
		}
	}
	return nil //for others
}

//Gets the SimplePokemon if the pokemon's name is equal to the key.
func GetByPokemonName(key string, pokemon Pokemon) map[string]SimplePokemon {
	pokemonMap := make(map[string]SimplePokemon)

	name := pokemon.Name
	pokemonMap[name] = NewSimplePokemon(pokemon)
	return pokemonMap
}