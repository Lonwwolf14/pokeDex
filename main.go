package main

import (
	"time"

	"example.com/pokedex/internal/pokeapi"
	"example.com/pokedex/internal/pokemonapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: *pokeClient,
		pokemonClient: *pokemonapi.NewClient(),
		pokedex:       make(map[string]pokemonapi.Pokemon),
	}
	startRepl(cfg)
}
