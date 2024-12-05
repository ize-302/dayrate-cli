package main

import (
	"flag"
	"fmt"

	"github.com/ize-302/dayrate-cli/internal/handlers"
	"github.com/ize-302/dayrate-cli/internal/utils"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		utils.GetHelp()
	} else if flag.Arg(0) == "help" {
		utils.GetHelp()
	} else if flag.Arg(0) == "add" {
		handlers.HandleRatingEntry()
	} else if flag.Arg(0) == "list" {
		handlers.HandleListRatigs()
	} else {
		fmt.Println("hello")
	}
}
