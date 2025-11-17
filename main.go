package main

import (
	"fmt"
	"os"
	"strconv"
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
			return
		}
		description := strings.Trim(args[1], "\"")

		if err := handlers.AddTask(description); err != nil {
			fmt.Println("Error:", err)
		}

	case "update":
		if len(args) < 3 {
			fmt.Printf("Usage: %s update <id> \"<description>\"\n", constants.PROGRAM_PREFIX)
			return
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		description := strings.Trim(args[2], "\"")

		if err := handlers.UpdateTask(id, description); err != nil {
			fmt.Println("Error:", err)
		}

	case "delete":
		if len(args) < 2 {
			fmt.Printf("Usage: %s delete <id>\n", constants.PROGRAM_PREFIX)
			return
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		if err := handlers.DeleteTask(id); err != nil {
			fmt.Println("Error:", err)
		}

	case "mark-todo":
		if len(args) < 2 {
			fmt.Printf("Usage: %s mark-todo <id>\n", constants.PROGRAM_PREFIX)
			return
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		if err := handlers.MarkTodo(id); err != nil {
			fmt.Println("Error:", err)
		}

	case "mark-in-progress":
		if len(args) < 2 {
			fmt.Printf("Usage: %s mark-in-progress <id>\n", constants.PROGRAM_PREFIX)
			return
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		if err := handlers.MarkInProgress(id); err != nil {
			fmt.Println("Error:", err)
		}

	case "list":
		if len(args) < 2 {
			fmt.Printf("Usage: %s list <todo/in-progress/done>\n", constants.PROGRAM_PREFIX)
			return
		}

		status := args[1]

		tasks, err := handlers.ListTasksByStatus(status)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		for _, t := range tasks {
			fmt.Printf("[%v] %-12s %s\n",
				t["id"],
				t["status"],
				t["description"],
			)
		}
	case "mark-done":
		if len(args) < 2 {
			fmt.Printf("Usage: %s mark-done <id>\n", constants.PROGRAM_PREFIX)
			return
		}

		id, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}

		if err := handlers.MarkDone(id); err != nil {
			fmt.Println("Error:", err)
		}
	case "help":
		handlers.HelpManual()

	default:
		fmt.Printf("task-cli %s: unknown command\n", args[0])
		fmt.Printf("Run '%s help' for usage.\n", constants.PROGRAM_PREFIX)
	}
}
