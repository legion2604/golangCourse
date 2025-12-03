package service

import (
	"fmt"
	"golangCourse/todo/models"
)

func DeleteTask(taskList *[]models.Task) {
	fmt.Println("Input task number to delete:")
	var num int
	fmt.Scanln(&num)
	if num > 0 && num <= len(*taskList) {
		*taskList = append((*taskList)[:num-1], (*taskList)[num:]...)
		fmt.Println("Task deleted!")
	}
}
