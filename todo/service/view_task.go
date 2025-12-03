package service

import (
	"fmt"
	"golangCourse/todo/models"
)

func ViewTasks(taskList []models.Task) {
	fmt.Println("Your Tasks:")
	for i, t := range taskList {
		fmt.Println(i+1, " Description:", t.Description, "| Status: ", t.Status, " | Created At: ", t.CreatedAt)
	}
}
