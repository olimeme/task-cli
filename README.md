# Task CLI

A simple command-line task management tool written in Go.  
Manage tasks directly from your terminal with add, update, delete, and status management commands. Tasks are stored in a local JSON file.

---

## Features

- Add, update, and delete tasks  
- Mark tasks as:
  - **To Do**
  - **In Progress**
  - **Done**
- List tasks by status or list all tasks  
- Local JSON storage (persistent)  
- Simple, fast, and lightweight

---

## Usage

```bash
task-cli <command> [arguments]
```
### Commands

| Command                                 | Description                           |
| --------------------------------------- | ------------------------------------- | 
| `add "<task description>"`              | Add a new task                        |
| `update <task ID> "<task description>"` | Update an existing task's description |
| `delete <task ID>`                      | Delete a task by ID                   |
| `mark-todo <task ID>`                   | Mark a task as "to do"                |
| `mark-in-progress <task ID>`            | Mark a task as "in progress"          |
| `mark-done <task ID>`                   | Mark a task as "done"                 |        
| `list [todo / in-progress /done]` | List tasks filtered by status, or all tasks if no argument is provided |
| `help`                                  | Show the help manual                  |      

## Data Storage

Tasks are stored locally in a JSON file named:

```
tasks.json
```

- The file is created automatically if it doesn’t exist.
- The CLI ensures data integrity and validates task fields.
- Corrupted files are automatically reset to prevent crashes.
