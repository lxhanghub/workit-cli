package main

import (
	"fmt"
	"os"

	"github.com/lxhanghub/mfish-cli/cmd"
)

func main() {
	fmt.Println(` 
 __  __ _____ _  ______ _    _
|  \/  |  ___| ||  ___|| |  | |
| |\/| | |__ | || |___ | |__| |
| |  | |  __|| ||___  || |--| |
|_|  |_|_|   |_||_____||_|  |_|
                               `)
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
