package main

import (
	"os"

	"github.com/lovehang/newb-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
