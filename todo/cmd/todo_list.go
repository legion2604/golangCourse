package main

import (
	"fmt"
	"golangCourse/todo/models"
	"golangCourse/todo/service"
)

func main() {
	fmt.Println("Hello it's your ToDo app!")
	taskList := make([]models.Task, 0)
	start(taskList)
}

func start(taskList []models.Task) string {
	fmt.Println("\n", "1. Add Task", "\n", "2. View Tasks", "\n", "3. Complete the task", "\n", "4. Delete Task", "\n", "5. Exit")
	var input int
	fmt.Scan(&input)

	switch input {
	case 1:
		service.AddTask(&taskList)
	case 2:
		service.ViewTasks(taskList)
	case 3:
		service.CompleteTask(&taskList)
	case 4:
		service.DeleteTask(&taskList)
	case 5:
		return "Bye!"
	default:
		fmt.Println("Invalid input")
	}

	return start(taskList)
}
