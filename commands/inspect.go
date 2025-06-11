package commands

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, name string) error {
	if name == "" {
		return errors.New("no name provided")
	}

	if pokemon, exists := cfg.Pokedex.Pokemon[name]; exists {
		fmt.Println("Name: ", pokemon.Name)
		fmt.Println("Height: ", pokemon.Height)
		fmt.Println("Stats: ")
		for _, stat := range pokemon.Stats {
			fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types: ")
		for _, t := range pokemon.Types {
			fmt.Println(" -", t.Type.Name)
		}
		return nil
	}

	return errors.New("you have not caught that pokemon")
}
