package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("explore takes exactly 1 argument")
	}
	location := args[0]

	locationAreaResp, err := cfg.pokeapiClient.GetLocationArea(location)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", locationAreaResp.Name)
	fmt.Println("Pokemon found:")
	for _, pokemon := range locationAreaResp.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
