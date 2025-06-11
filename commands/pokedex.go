package commands

import (
	"fmt"
)

func commandPokedex(cfg *Config, _ string) error {

	fmt.Println("Your pokedex:")
	for _, pokemon := range cfg.Pokedex.Pokemon {
		fmt.Println("-", pokemon.Name)
	}

	return nil
}
