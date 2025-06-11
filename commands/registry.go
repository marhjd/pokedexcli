package commands

import (
	"errors"

	"github.com/marhjd/pokedexcli/internal/pokecache"
)

type CLICommand struct {
	Name        string
	Description string
	Callback    func(cfg *Config, s string) error
}

type Config struct {
	Next     string
	Previous string
	Cache    pokecache.Cache
	Pokedex  struct {
		Pokemon map[string]Pokemon
	}
}

const baseURL = "https://pokeapi.co/api/v2/"

func (cfg *Config) GetNext() string {
	if cfg.Next == "" {
		return baseURL + "location-area/"
	}
	return cfg.Next
}

func (cfg *Config) GetPrevious() (string, error) {
	if cfg.Previous == "" {
		return "", errors.New("you're on the first page")
	}
	return cfg.Previous, nil
}

func GetSupportedCommands() map[string]CLICommand {
	supportedCommands := map[string]CLICommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "List the next 20 locations",
			Callback:    commandMap,
		},
		"mapb": {
			Name:        "mapb",
			Description: "List the previous 20 locations",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore",
			Description: "Explore the pokemon in an area",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch",
			Description: "Catch a pokemon",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "Inspect a pokemon",
			Callback:    commandInspect,
		},
	}
	return supportedCommands
}
