package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/marhjd/pokedexcli/commands"
	"github.com/marhjd/pokedexcli/internal/pokecache"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	fields := strings.Fields(lowerText)
	return fields
}

func StartRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	cfg := new(commands.Config)
	cfg.Cache = *pokecache.NewCache(1 * time.Minute)
	cfg.Pokedex.Pokemon = make(map[string]commands.Pokemon)
	for {
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			userInput := scanner.Text()
			words := strings.Fields(strings.ToLower(userInput))
			if len(words) > 0 {
				userCmd := words[0]
				userParam := ""
				if len(words) > 1 {
					userParam = words[1]
				}
				supportedCmds := commands.GetSupportedCommands()
				if cmd, exists := supportedCmds[userCmd]; exists {
					if err := cmd.Callback(cfg, userParam); err != nil {
						fmt.Println(err)
					}
				} else {
					fmt.Println("Unknown command")
				}
			}
		}
	}
}
