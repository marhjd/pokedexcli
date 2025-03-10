package commands

type CLICommand struct {
	Name        string
	Description string
	Callback    func() error
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
	}
	return supportedCommands
}
