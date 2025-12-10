package service

import (
	"errors"
	"fmt"
	"golangCourse/todo/models"
)

func DeleteTask(taskList *[]models.Task) error {
	fmt.Println("Input task number to delete:")
	var num int
	fmt.Scan(&num)
	if num > len(*taskList) {
		err := errors.New("Input task number to delete is greater than the number of tasks")
		return err
	}
	if num > 0 && num <= len(*taskList) {
		*taskList = append((*taskList)[:num-1], (*taskList)[num:]...)
		fmt.Println("Task deleted!")
	}
	return nil
}
