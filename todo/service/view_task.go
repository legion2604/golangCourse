package service

import (
	"errors"
	"fmt"
	"golangCourse/todo/models"
	"log"
)

func ViewTasks(taskList []models.Task) error {
	if len(taskList) == 0 {
		err := errors.New("no tasks available")
		log.Println(err)
		return err
	}
	fmt.Println("Your Tasks:")
	for i, t := range taskList {
		fmt.Println(i+1, " Description:", t.Description, "| Status: ", t.Status, " | Created At: ", t.CreatedAt)
	}
	log.Println("успешный показ задачи")
	return nil
}
