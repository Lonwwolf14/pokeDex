package main

import (
	"time"

	"example.com/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	cfg := &config{
		pokeapiClient: *pokeClient,
	}
	startRepl(cfg)
}
