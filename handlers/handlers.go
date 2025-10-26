package handlers

import (
	"fmt"

	"github.com/olimeme/constants"
)

func HelpManual() {
	fmt.Printf("Usage:\n\n\t%s <command> \n\n", constants.PROGRAM_PREFIX)

	fmt.Println("Commands:\n")
	fmt.Printf("\t%-40s %-10s\n", "add \"[task description]\":", "add a task to the task list")
	fmt.Printf("\t%-40s %-10s\n", "update [task ID] \"[task description]\":", "update the task description")
	fmt.Printf("\t%-40s %-10s\n", "delete [task ID]:", "delete the task")
	fmt.Printf("\t%-40s %-10s\n", "mark-todo [task ID]:", "mark task to do")
	fmt.Printf("\t%-40s %-10s\n", "mark-in-progress [task ID]:", "mark task done")
	fmt.Printf("\t%-40s %-10s\n", "list [todo/in-progress/done]:", "list all the tasks to do/in-progress/done")
}
