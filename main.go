package main

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/nicolas-cuveillier/gocheck/core/cli"
)

var author = "Nicolas Cuveillier"
var version = "0.0.0"

func showBanner() {
	name := fmt.Sprintf("gc2 (v%s)", version)
	banner := `
_________   ___________ _________________  _________ ______
__  ____/_______  ____/   / __  /   ___// ____/ /  /_/  _/
_  / __ _  __ \_/ /   /  /__/  //  __  /  /  __/    ___/
/ /_/ / / /_/ // /__ /  /--/  //  ___ /  /___ /  /\  \
\____/  \____/\____/ \_/  /__/ \____/ \_____//__/  \__\
											
	`

	all_lines := strings.Split(banner, "\n")
	w := len(all_lines[1])

	fmt.Println(banner)

	color.Green(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(name))/2, name)))
	color.Blue(fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(author))/2, author)))

	fmt.Println()
}

func main() {

	showBanner()

	p := prompt.New(
		cli.CoreExecutor,
		cli.CoreCompleter,
		prompt.OptionPrefix("[gocheck] > "),
	)
	p.Run()
}
