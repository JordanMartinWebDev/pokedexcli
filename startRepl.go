package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jordanmartinwebdev/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	input := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex >")
		input.Scan()
		words := cleanInput(input.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"map": {
			name:        "map",
			description: "Displays page of 20 Poke world locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous page of 20 Poke world locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Displays pokemon in an area",
			callback:    commandExplore,
		},
		"inspect": {
			name:        "inspect {pokemon_name(caught only)}",
			description: "Displays Pokemon Information",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all caught pokemon",
			callback:    commandPokedex,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Chance to catch the listed Pokemon",
			callback:    commandCatch,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}
