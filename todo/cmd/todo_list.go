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

	if input == 1 {
		service.AddTask(&taskList)
	} else if input == 2 {
		service.ViewTasks(taskList)
	} else if input == 3 {
		service.CompleteTask(&taskList)
	} else if input == 4 {
		service.DeleteTask(&taskList)
	} else if input == 5 {
		return "Bye!"
	} else {
		fmt.Println("Invalid input")
	}

	return start(taskList)
}
