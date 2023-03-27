package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

// list all tasks from the array with an icon
func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas.")
		return
	}

	for _, task := range tasks {
		status := " "
		if task.Complete {
			status = "✓"
		}
		fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Name)
	}
}

// add a new task to the array
func AddTask(tasks []Task, name string) []Task {
	newTask := Task{
		ID:       GetNextID(tasks),
		Name:     name,
		Complete: false,
	}
	return append(tasks, newTask)
}

// search task with id and mark as completed
