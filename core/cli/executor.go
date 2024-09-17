package cli

import (
	"fmt"
	"os"

	"github.com/nicolas-cuveillier/gocheck/core/utils"
)

func CoreExecutor(s string) {
	cmd, args := utils.ParseCommand(s)

	switch cmd {
	case "help":
		CmdHelp()
	case "generate":
		CmdGenerate(args)
	case "crack":
		CmdCrack(args)
	case "check":
		CmdCheck(args)
	case "exit":
		os.Exit(0)
		return
	case "":
	default:
		return
	}
}

func CmdHelp() {
	fmt.Println()

	fmt.Printf("%-15s %s \n", "help", "Show the list of available commands")

	fmt.Printf("%-15s %s \n", "generate", "Generate a new password according to the parameters provided")
	fmt.Printf("    %-10s %s \n", "-l <number>", "Provide the length of the password (default 20)")
	fmt.Printf("    %-11s %s \n", "-c", "Use lowercase characters to generate the password")
	fmt.Printf("    %-11s %s \n", "-u", "Use uppercase characters to generate the password")
	fmt.Printf("    %-11s %s \n", "-s", "Use special characters to generate the password")

	fmt.Printf("%-15s %s \n", "check", "Check the strength of the password")

	fmt.Printf("%-15s %s \n", "crack", "Crack the password")
	fmt.Printf("    %-11s %s \n", "-b", "Use brute-force to crack the password")
	fmt.Printf("    %-10s %s \n", "-l <number>", "Give the max length for the brute-force attack")
	fmt.Printf("    %-11s %s \n", "-d", "Use dictionary to crack the password")

	fmt.Printf("%-15s %s \n", "exit", "Exit the program")

	fmt.Println()
}
