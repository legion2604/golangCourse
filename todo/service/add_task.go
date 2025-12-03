package service

import (
	"fmt"
	"golangCourse/todo/models"
	"golangCourse/todo/utils"
	"time"
)

const (
	layoutISO = "Jan 2 15:04:05"
)

func AddTask(taskList *[]models.Task) {
	fmt.Println("Input task description:")
	desc := utils.Print()
	date := time.Now().Format(layoutISO)
	newTask := models.Task{Description: desc, Status: false, CreatedAt: date}
	*taskList = append(*taskList, newTask)
	fmt.Println("Task added!")
}
