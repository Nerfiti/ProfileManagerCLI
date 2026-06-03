package main

import (
	"os"

	"mws/commands"
)

func main() {
	if err := commands.Execute(); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}
