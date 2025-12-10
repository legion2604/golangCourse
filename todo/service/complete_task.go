package service

import (
	"errors"
	"fmt"
	"golangCourse/todo/models"
)

func CompleteTask(taskList *[]models.Task) error {

	fmt.Println("Input task number to complete:")
	var num int
	fmt.Scanln(&num)
	if num > 0 && num <= len(*taskList) {
		(*taskList)[num-1].Status = true
		fmt.Println("Task marked as completed!")
	} else {
		err := errors.New("invalid input")
		return err
	}
	return nil
}
