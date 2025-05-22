package commands

import (
	"fmt"
)

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range GetSupportedCommands() {
		fmt.Printf("%v: %v\n", command.Name, command.Description)
	}

	return nil
}
