package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/olimeme/constants"
)

func HelpManual() {
	fmt.Printf("Usage:\n\n\t%s <command> \n\n", constants.PROGRAM_PREFIX)

	fmt.Println("Commands:\n")
	fmt.Printf("\t%-40s %-10s\n", "add \"<task description>\":", "add a task to the task list")
	fmt.Printf("\t%-40s %-10s\n", "update <task ID> \"<task description>\":", "update the task description")
	fmt.Printf("\t%-40s %-10s\n", "delete <task ID>:", "delete the task")
	fmt.Printf("\t%-40s %-10s\n", "mark-todo <task ID>:", "mark task to do")
	fmt.Printf("\t%-40s %-10s\n", "mark-in-progress <task ID>:", "mark task done")
	fmt.Printf("\t%-40s %-10s\n", "list <todo/in-progress/done>:", "list all the tasks to do/in-progress/done")
}

func AddTask(description string) error {
	jsonFile, err := os.OpenFile(constants.FILENAME, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	data := map[string][]map[string]interface{}{
		"tasks": {},
	}

	if len(byteValue) > 0 {
		err = json.Unmarshal(byteValue, &data)
		if err != nil {
			return err
		}
	}

	newTask := map[string]interface{}{}
	newTask["id"] = len(data["tasks"]) + 1
	newTask["description"] = description
	newTask["status"] = "to do"
	newTask["createdAt"] = time.Now()
	newTask["updatedAt"] = time.Now()

	data["tasks"] = append(data["tasks"], newTask)

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(constants.FILENAME, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func UpdateTask(id int, description string) error {
	return nil
}

func DeleteTask(id int) error {
	return nil
}
