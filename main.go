package main

import (
	"os"

	"github.com/georgetaylor/spotctl/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
