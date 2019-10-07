package main

import (
	"fmt"
	"os"

	"github.com/Pandentia/pride/pride"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	app := kingpin.New("pride", "A simple executable for printing pride flags")

	// register the commands
	flagCommands := make(map[string]*kingpin.CmdClause)
	for flag := range pride.Flags {
		flagCommands[flag] = app.Command(flag, fmt.Sprintf("A %s pride flag", flag))
	}

	// parse arguments
	parseResult := kingpin.MustParse(app.Parse(os.Args[1:]))

	// print the flag that was asked for
	for flag, flagRunner := range pride.Flags {
		if flagCommands[flag].FullCommand() == parseResult {
			flagRunner()
		}
	}
}
