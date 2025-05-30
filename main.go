package main

import (
	"os"

	"github.com/georgetaylor/rackspace-spot-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
