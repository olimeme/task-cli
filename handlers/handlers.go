package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/olimeme/constants"
)

func loadTasks() (map[string][]map[string]interface{}, error) {
	jsonFile, err := os.OpenFile(constants.FILENAME, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	data := map[string][]map[string]interface{}{
		"tasks": {},
	}

	if len(byteValue) > 0 {
		if err := json.Unmarshal(byteValue, &data); err != nil {
			return nil, err
		}
	}

	return data, nil
}

func saveTasks(data map[string][]map[string]interface{}) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(constants.FILENAME, jsonData, 0644)
}

func updateTaskStatus(id int, status string) error {
	data, err := loadTasks()
	if err != nil {
		return err
	}

	found := false
	for _, task := range data["tasks"] {
		if int(task["id"].(float64)) == id {
			task["status"] = status
			task["updatedAt"] = time.Now()
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return saveTasks(data)
}

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
	data, err := loadTasks()
	if err != nil {
		return err
	}

	newTask := map[string]interface{}{
		"id":          len(data["tasks"]) + 1,
		"description": description,
		"status":      "to do",
		"createdAt":   time.Now(),
		"updatedAt":   time.Now(),
	}

	data["tasks"] = append(data["tasks"], newTask)

	return saveTasks(data)
}

func UpdateTask(id int, description string) error {
	data, err := loadTasks()
	if err != nil {
		return err
	}

	found := false
	for _, task := range data["tasks"] {
		if int(task["id"].(float64)) == id {
			task["description"] = description
			task["updatedAt"] = time.Now()
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	return saveTasks(data)
}

func DeleteTask(id int) error {
	data, err := loadTasks()
	if err != nil {
		return err
	}

	index := -1
	for i, task := range data["tasks"] {
		if int(task["id"].(float64)) == id {
			index = i
			break
		}
	}

	if index == -1 {
		return fmt.Errorf("task with ID %d not found", id)
	}

	data["tasks"] = append(data["tasks"][:index], data["tasks"][index+1:]...)

	return saveTasks(data)
}

func ListTasksByStatus(status string) ([]map[string]interface{}, error) {
	valid := map[string]bool{
		"todo":        true,
		"in-progress": true,
		"done":        true,
	}

	if !valid[status] {
		return nil, fmt.Errorf("invalid list argument: %s (allowed: todo, in-progress, done)", status)
	}

	jsonFile, err := os.Open(constants.FILENAME)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	data := map[string][]map[string]interface{}{
		"tasks": {},
	}

	if len(byteValue) > 0 {
		if err = json.Unmarshal(byteValue, &data); err != nil {
			return nil, err
		}
	}

	filtered := []map[string]interface{}{}
	for _, task := range data["tasks"] {
		taskStatus := task["status"].(string)

		if status == "todo" && taskStatus == "to do" {
			filtered = append(filtered, task)
		}

		if status == "in-progress" && taskStatus == "in progress" {
			filtered = append(filtered, task)
		}

		if status == "done" && taskStatus == "done" {
			filtered = append(filtered, task)
		}
	}

	return filtered, nil
}

func MarkTodo(id int) error {
	return updateTaskStatus(id, "to do")
}

func MarkInProgress(id int) error {
	return updateTaskStatus(id, "in progress")
}

func MarkDone(id int) error {
	return updateTaskStatus(id, "done")
}
