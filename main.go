package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/marhjd/pokedexcli/commands"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			words := strings.Fields(strings.ToLower(userInput))
			if len(words) > 0 {
				userCmd := words[0]
				supportedCmds := commands.GetSupportedCommands()
				if cmd, exists := supportedCmds[userCmd]; exists {
					if err := cmd.Callback(); err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println("Unknown command")
				}
			}
		}
	}
}
