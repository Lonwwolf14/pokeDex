package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/pokedex/internal/pokeapi"
	"example.com/pokedex/internal/pokemonapi"
)

type config struct {
	pokeapiClient   pokeapi.Client
	nextLocationUrl *string
	prevLocationUrl *string
	pokemonClient   pokemonapi.Client
	pokedex         map[string]pokemonapi.Pokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		commandName := words[0]

		args := words[1:]
		command, exists := getCommands()[commandName]
		if !exists {
			fmt.Println("Invalid command")
			continue
		} else {
			err := command.callBack(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		}

	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	output = strings.TrimSpace(output)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name        string
	description string
	callBack    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callBack:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the current location",
			callBack:    commandMapF,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous location",
			callBack:    commandMapB,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callBack:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Explore the area",
			callBack:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a pokemon",
			callBack:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a pokemon",
			callBack:    commandInspect,
		},
	}
}
