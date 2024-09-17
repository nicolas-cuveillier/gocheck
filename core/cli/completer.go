package cli

import (
	"github.com/c-bata/go-prompt"
)

var coreCommands = []prompt.Suggest{
	{Text: "help", Description: "Show the list of available commands"},
	{Text: "check", Description: "Check the strength of a password"},
	{Text: "generate", Description: "Generate a password"},
	{Text: "crack", Description: "Crack the provided password"},
	{Text: "exit", Description: "Exit the program"},
}

func CoreCompleter(d prompt.Document) []prompt.Suggest {
	return coreCommands
}
