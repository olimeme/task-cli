package main

import (
	"fmt"
	"os"

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
		fmt.Println("Adding task...")
	case "update":
		fmt.Println("Updating task...")
	case "delete":
		fmt.Println("Deleting task...")
	case "help":
		handlers.HelpManual()
	default:
		fmt.Printf("task-cli %s: unknown command\n", args[0])
		fmt.Printf("Run '%s help' for usage.", constants.PROGRAM_PREFIX)
	}
}
