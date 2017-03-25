package main

import (
	"github.com/gorilla/mux"
	"strings"
	"gopkg.in/yaml.v2"
	"fmt"
	"net/http"
	"net/url"
)

func ListHandler(w http.ResponseWriter, r *http.Request) {
	s := r.URL.String()
	u, err := url.Parse(s)
	Check(err)
	var str string
	if u.RawQuery != "" { //if URL contains query
		m, _ := url.ParseQuery(u.RawQuery) //map for query
		pokemons := pokemonBD
		for key := range m {
			kk := strings.ToLower(string(key))
			value := strings.ToLower(m[key][0])
			switch kk {
			case "type": //if the query key is "type", lists all pokemons whose types are equal query value.
				pokemons = FilterByTypeName(pokemons, value)
				break
			case "move": //if the query key is "move", lists all pokemons whose even one of attacks(specialAttacks or fastAttacks) is equal query value.
				pokemons = FilterByMoveName(pokemons, value)
				break
			case "sortby": //if the query key is "sortby", lists pokemons by their sortable pokemon property.
				pokemons = SortByPokemonProperty(pokemons, value)
				break
			default:
				pokemons = make([]Pokemon, 0)
			}

		}
		WriteSimplePokemons(w, NewSimplePokemons(pokemons));
	} else { //if URL does not contain query
		vars := mux.Vars(r)
		key := strings.ToLower(vars["key"])

		var obj interface{}

		switch key {
		case "", "pokemons": //if there is not a key (http://localhost:8080/list) or the key is "pokemons", lists all pokemons.
			WriteSimplePokemons(w, NewSimplePokemons(pokemonBD))
			return
		case "types": //if the key is "types", lists all types.
			obj = typeBD
			break
		case "moves": //if the key is "moves", lists all moves.
			obj = moveBD
			break
		}
		d, _ := yaml.Marshal(&obj)
		str = string(d)
		fmt.Fprint(w, str)
	}
}

func WriteSimplePokemon(w http.ResponseWriter, pokemon SimplePokemon) {
	toWrite := make(map[string]SimplePokemon)
	toWrite[pokemon.Name] = pokemon
	dat, err := yaml.Marshal(toWrite)
	Check(err)
	fmt.Fprint(w, string(dat))
}

func WriteSimplePokemons(w http.ResponseWriter, pokemons []SimplePokemon) {
	for _, pokemon := range pokemons {
		WriteSimplePokemon(w, pokemon)
	}
}

//Lists all pokemons, if the pokemonsToFilter' TypeI or TypeII is equal to the typeName.
func FilterByTypeName(pokemonsToFilter []Pokemon, typeName string) []Pokemon {
	pokemons := make([]Pokemon, 0)
	for _, pokemon := range pokemonsToFilter {
		var b bool = false
		for _, y := range pokemon.TypeI {
			if typeName == strings.ToLower(y) {
				b = true
				pokemons = append(pokemons, pokemon)
				break
			}
		}
		if b == false {
			for _, y := range pokemon.TypeII {
				if typeName == strings.ToLower(y) {
					b = true
					pokemons = append(pokemons, pokemon)
					break
				}
			}
		}
	}
	return pokemons
}

//Lists all pokemons, if the pokemons' even one of attacks(specialAttacks or fastAttacks) is equal to the moveName.
func FilterByMoveName(pokemonsToFilter []Pokemon, moveName string) []Pokemon {
	pokemons := make([]Pokemon, 0)

	for _, pokemon := range pokemonsToFilter {
		var b bool = false
		for _, move := range pokemon.SpecialAttacks {
			if moveName == strings.ToLower(move) {
				b = true
				pokemons = append(pokemons, pokemon)
			}
		}

		if b == false {
			for _, move := range pokemon.FastAttackS {
				if moveName == strings.ToLower(move) {
					b = true
					pokemons = append(pokemons, pokemon)
				}
			}
		}
	}

	return pokemons
}

//Sorts pokemonsToList by the key.
func SortByPokemonProperty(pokemonsToList []Pokemon, key string) []Pokemon {

	pokemonNameMap := make(map[string]int)
	sortedPokemonNames := make(map[string]int)
	sortedPokemonBD := make([]Pokemon, 0)
	var b bool = false
	if key == strings.ToLower("BaseAttack") { //If pokemons are going to be sorted by their BaseAttack
		b = true
		for _, pokemon := range pokemonsToList {
			pokemonNameMap[pokemon.Name] = pokemon.BaseAttack
		}
	}
	if b != true {
		if key == strings.ToLower("BaseDefense") { //If pokemons are going to be sorted by their BaseDefense
			b = true
			for _, pokemon := range pokemonsToList {
				pokemonNameMap[pokemon.Name] = pokemon.BaseDefense
			}
		}
	}
	if b != true {
		if key == strings.ToLower("BaseStamina") { //If pokemons are going to be sorted by their BaseStamina
			for _, pokemon := range pokemonsToList {
				pokemonNameMap[pokemon.Name] = pokemon.BaseStamina
			}
		}
	}
	sortedPokemonNames = SortMap(pokemonNameMap) //pokemonNameBaseAttack is sorted
	sortedPokemonBD = make([]Pokemon, len(pokemonsToList))
	for _, pokemon := range pokemonsToList {
		sortedPokemonBD[sortedPokemonNames[pokemon.Name]-1] = pokemon
	}
	return sortedPokemonBD
}