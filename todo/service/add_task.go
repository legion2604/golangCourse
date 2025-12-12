package service

import (
	"errors"
	"fmt"
	"golangCourse/todo/models"
	"golangCourse/todo/utils"
	"time"
)

const (
	layoutISO = "Jan 2 15:04:05"
)

func AddTask(taskList *[]models.Task) error {
	fmt.Println("Input task description:")
	desc := utils.Print()
	if desc == "" {
		err := errors.New("invalid input")
		return err
	}
	date := time.Now().Format(layoutISO)
	newTask := models.Task{Description: desc, Status: false, CreatedAt: date}
	*taskList = append(*taskList, newTask)
	return nil
}
