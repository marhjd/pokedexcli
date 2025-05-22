package commands

type CLICommand struct {
	Name        string
	Description string
	Callback    func(cfg *Config) error
}

type Config struct {
	Next     string
	Previous string
}

func (cfg *Config) GetNext() string {
	if cfg.Next == "" {
		return "https://pokeapi.co/api/v2/location-area/"
	}
	return cfg.Next
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
	}
	return supportedCommands
}
