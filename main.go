package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/olimeme/constants"
	"github.com/olimeme/handlers"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Usage: %s <command>", constants.PROGRAM_PREFIX)
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 2 {
			fmt.Printf("Usage: %s add \"<description>\"\n", constants.PROGRAM_PREFIX)
		}
		description := strings.Trim(args[1], "\"")

		err := handlers.AddTask(description)

		if err != nil {
			fmt.Println(err)
		}
	case "update":
		//handlers.UpdateTask(id, description)
	case "delete":
		//handlers.DeleteTask(id)
	case "help":
		handlers.HelpManual()
	default:
		fmt.Printf("task-cli %s: unknown command\n", args[0])
		fmt.Printf("Run '%s help' for usage.", constants.PROGRAM_PREFIX)
	}
}
