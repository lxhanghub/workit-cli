package main

import (
	"os"

	"github.com/lxhanghub/mfish-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
