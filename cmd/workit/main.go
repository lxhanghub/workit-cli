package main

import (
	"fmt"
	"os"

	"github.com/lxhanghub/workit-cli/cmd"
)

func main() {
	fmt.Println(` 
__        __  _______   _ ____    _  ____   _   _______
\ \      / / /  ___  \ | |   ))  | |/ ___| | | |__   __|
 \ \ /\ / /  | |   | | | |__))   | | |     | |    | |
  \ V  V /   | |___| | | |_ \__  | | |___  | |    | |
   \_/\_/    \_______/ |_| \___| |_|\____| |_|    |_|
                               `)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
