package main

import (
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("can Catch Exaclty 1 Pokemon")
	}
	pokemonName := args[0]

	pokemonResp, err := cfg.pokemonClient.GetPokemon(pokemonName)
	if err != nil {
		return fmt.Errorf("pokemon not found")
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	rand.Seed(time.Now().UnixNano())
	if rand.Intn(300) < (300 - pokemonResp.BaseExperience) {
		fmt.Printf("%s was caught!\n", pokemonName)
		fmt.Println("You may now inspect it with the inspect command.")
		cfg.pokedex[pokemonName] = pokemonResp
	} else {
		fmt.Printf("%s escaped!\n", pokemonName)
	}
	return nil
}
