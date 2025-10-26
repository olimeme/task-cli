package main

import (
	"fmt"
	"os"
)

const PROGRAM_PREFIX = "task-cli"

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Printf("Usage: %s <command>", PROGRAM_PREFIX)
		return
	}

	switch args[0] {
	case "add":
		fmt.Println("Adding task...")
	case "update":
		fmt.Println("Updating task...")
	case "delete":
		fmt.Println("Deleting task...")
	default:
		fmt.Println("Unknown command:", args[0])
	}
}
