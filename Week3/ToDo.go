package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Task represents a single to-do item
type Task struct {
	ID    int
	Title string
	Done  bool
}

// TaskManager manages the list of tasks
type TaskManager struct {
	Tasks  []Task
	NextID int
}

// AddTask adds a new task to the list
func (tm *TaskManager) AddTask(title string) {
	tm.Tasks = append(tm.Tasks, Task{ID: tm.NextID, Title: title, Done: false})
	tm.NextID++
	fmt.Println("Task added successfully!")
}

// DeleteTask removes a task by ID
func (tm *TaskManager) DeleteTask(id int) {
	for i, task := range tm.Tasks {
		if task.ID == id {
			tm.Tasks = append(tm.Tasks[:i], tm.Tasks[i+1:]...)
			fmt.Println("Task deleted successfully!")
			return
		}
	}
	fmt.Println("Task not found.")
}

// MarkTaskDone marks a task as done by ID
func (tm *TaskManager) MarkTaskDone(id int) {
	for i := range tm.Tasks {
		if tm.Tasks[i].ID == id {
			tm.Tasks[i].Done = true
			fmt.Println("Task marked as done!")
			return
		}
	}
	fmt.Println("Task not found.")
}

// ListTasks lists all tasks
func (tm *TaskManager) ListTasks() {
	if len(tm.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("\nTo-Do List:")
	for _, task := range tm.Tasks {
		status := "Pending"
		if task.Done {
			status = "Done"
		}
		fmt.Printf("[%d] %s - %s\n", task.ID, task.Title, status)
	}
	fmt.Println()
}

// Main program loop
func main() {
	tm := TaskManager{}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("To-Do List App")
		fmt.Println("1. Add Task")
		fmt.Println("2. Delete Task")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. List Tasks")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Enter task title: ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())
			tm.AddTask(title)
		case "2":
			fmt.Print("Enter task ID to delete: ")
			scanner.Scan()
			var id int
			fmt.Sscanf(scanner.Text(), "%d", &id)
			tm.DeleteTask(id)
		case "3":
			fmt.Print("Enter task ID to mark as done: ")
			scanner.Scan()
			var id int
			fmt.Sscanf(scanner.Text(), "%d", &id)
			tm.MarkTaskDone(id)
		case "4":
			tm.ListTasks()
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		fmt.Println()
	}
}
