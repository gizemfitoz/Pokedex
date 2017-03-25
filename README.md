# Pokedex

A Pokedex application in Go.

Run these commands before to run.
```
go get github.com/gorilla/mux gopkg.in/yaml.v2
```

This application handles requests for these routes:
- http://localhost:8080/{type-name}
- http://localhost:8080/{move-name}
- http://localhost:8080/list/types
- http://localhost:8080/list/pokemons
- http://localhost:8080/list/moves
- http://localhost:8080/list?type={type-name}
- http://localhost:8080/list?move={move-name}
- http://localhost:8080/{pokemon-name}
- http://localhost:8080/get/{pokemon-name}
- http://localhost:8080/get/{type-name}
- http://localhost:8080/get/{move-name}
- http://localhost:8080/get?name={pokemon-name}
- http://localhost:8080/get?name={type-name}
- http://localhost:8080/get?name={move-name}
- http://localhost:8080/list?sortby={sortable-pokemon-property}
- http://localhost:8080/list?type={type-name}&sortby={sortable-pokemon-property}
