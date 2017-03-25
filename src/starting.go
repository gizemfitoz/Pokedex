package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

var baseData BaseData
var typeBD []Type
var pokemonBD []Pokemon
var moveBD []Move

type Type struct {
	// Name of the type
	Name string `json:"name"`
	// The effective types, damage multiplize 2x
	EffectiveAgainst []string `json:"effectiveAgainst"`
	// The weak types that against, damage multiplize 0.5x
	WeakAgainst []string `json:"weakAgainst"`
}

type Pokemon struct {
	Number         string   `json:"Number"`
	Name           string   `json:"Name"`
	Classification string   `json:"Classification"`
	TypeI          []string `json:"Type I"`
	TypeII         []string `json:"Type II,omitempty"`
	Weaknesses     []string `json:"Weaknesses"`
	FastAttackS    []string `json:"Fast Attack(s)"`
	Weight         string   `json:"Weight"`
	Height         string   `json:"Height"`
	Candy struct {
		Name     string `json:"Name"`
		FamilyID int    `json:"FamilyID"`
	} `json:"Candy"`
	NextEvolutionRequirements struct {
		Amount int    `json:"Amount"`
		Family int    `json:"Family"`
		Name   string `json:"Name"`
	} `json:"Next Evolution Requirements,omitempty"`
	NextEvolutions []struct {
		Number string `json:"Number"`
		Name   string `json:"Name"`
	} `json:"Next evolution(s),omitempty"`
	PreviousEvolutions []struct {
		Number string `json:"Number"`
		Name   string `json:"Name"`
	} `json:"Previous evolution(s),omitempty"`
	SpecialAttacks      []string `json:"Special Attack(s)"`
	BaseAttack          int      `json:"BaseAttack"`
	BaseDefense         int      `json:"BaseDefense"`
	BaseStamina         int      `json:"BaseStamina"`
	CaptureRate         float64  `json:"CaptureRate"`
	FleeRate            float64  `json:"FleeRate"`
	BuddyDistanceNeeded int      `json:"BuddyDistanceNeeded"`
}

type Move struct {
	// The ID of the move
	ID int `json:"id"`
	// Name of the attack
	Name string `json:"name"`
	// Type of attack
	Type string `json:"type"`
	// The damage that enemy will take
	Damage int `json:"damage"`
	// Energy requirement of the attack
	Energy int `json:"energy"`
	// Dps is Damage Per Second
	Dps float64 `json:"dps"`
	// The duration
	Duration int `json:"duration"`
}

// BaseData is a struct for reading data.json
type BaseData struct {
	Types    []Type    `json:"types"`
	Pokemons []Pokemon `json:"pokemons"`
	Moves    []Move    `json:"moves"`
}

type SimplePokemon struct {
	Number         string   `json:"Number"`
	Name		   string	`json:"Name"`
	Weight         string	`json:"Weight"`
	Height         string	`json:"Height"`
	TypeI          []string `json:"Type I"`
	TypeII         []string `json:"Type II,omitempty"`
	Classification string   `json:"Classification"`
	BaseAttack     int		`json:"BaseAttack"`
	BaseDefense    int		`json:"BaseDefense"`
	BaseStamina    int		`json:"BaseStamina"`
	SpecialAttacks []string	`json:"Special Attack(s)"`
	FastAttacks    []string	`json:"Fast Attack(s)"`
	PreviousEvolutions []struct {
		Number string `json:"Number"`
		Name   string `json:"Name"`
	} `json:"Previous evolution(s),omitempty"`
	NextEvolutions []struct {
		Number string `json:"Number"`
		Name   string `json:"Name"`
	} `json:"Next evolution(s),omitempty"`
}

func NewSimplePokemon(pokemon Pokemon) SimplePokemon {
	simple := SimplePokemon{}
	simple.Number = pokemon.Number
	simple.Name = pokemon.Name
	simple.Weight = pokemon.Weight
	simple.Height = pokemon.Height
	simple.TypeI = pokemon.TypeI
	simple.TypeII = pokemon.TypeII
	simple.Classification = pokemon.Classification
	simple.BaseAttack = pokemon.BaseAttack
	simple.BaseDefense = pokemon.BaseDefense
	simple.BaseStamina = pokemon.BaseStamina
	simple.SpecialAttacks = pokemon.SpecialAttacks
	simple.FastAttacks = pokemon.FastAttackS
	simple.PreviousEvolutions = pokemon.PreviousEvolutions
	simple.NextEvolutions = pokemon.NextEvolutions
	return simple
}

func NewSimplePokemons(pokemons []Pokemon) []SimplePokemon {
	simplePokemons := make([]SimplePokemon, len(pokemons))
	for i, pokemon := range pokemons {
		simplePokemons[i] = NewSimplePokemon(pokemon)
	}
	return simplePokemons
}

//It is going to be use to sort map.
func SortMap(m map[string]int) map[string]int {
	n := map[int][]string{}
	var a []int
	for k, v := range m {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	keysAndIndexes := make(map[string]int)

	index := 0
	for _, k := range a {
		for _, key := range n[k] {
			index++
			keysAndIndexes[key] = index
		}
	}
	return keysAndIndexes
}

func Check(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func main() {
	data, err := ioutil.ReadFile("./src/data.json")
	Check(err)

	json.Unmarshal(data, &baseData)
	typeBD = baseData.Types
	pokemonBD = baseData.Pokemons
	moveBD = baseData.Moves

	router := mux.NewRouter()
	router.HandleFunc("/list", ListHandler)
	router.HandleFunc("/get", GetHandler)
	router.HandleFunc("/", GetByNameOrListByTypeOrMove)
	router.HandleFunc("/{key}", GetByNameOrListByTypeOrMove)
	router.HandleFunc("/list/{key}", ListHandler)
	router.HandleFunc("/get/{key}", GetHandler)

	log.Println("starting server on :8080")
	http.ListenAndServe(":8080", router)
}
