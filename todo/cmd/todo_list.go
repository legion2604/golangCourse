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
		err := service.AddTask(&taskList)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Task added!")
		}
	case 2:
		err := service.ViewTasks(taskList)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("End of task list.")
		}
	case 3:
		err := service.CompleteTask(&taskList)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Task complete!")
		}
	case 4:
		err := service.DeleteTask(&taskList)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Deleted successfully!")
		}
	case 5:
		return "Bye!"
	default:
		fmt.Println("Invalid input")
	}
	return start(taskList)
}
