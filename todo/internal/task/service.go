package task

import (
	"errors"
	"fmt"
	"golangCourse/todo/utils"
	"log"
	"time"
)

const (
	layoutISO = "Jan 2 15:04:05"
)

func AddTask(taskList *[]Task) error {
	fmt.Println("Input task description:")
	desc := utils.Print()
	if desc == "" {
		err := errors.New("invalid input")
		log.Println(err)
		return err
	}
	date := time.Now().Format(layoutISO)
	newTask := Task{Description: desc, Status: false, CreatedAt: date}
	*taskList = append(*taskList, newTask)
	log.Println("Задача успешно добавлена")
	return nil
}

func CompleteTask(taskList *[]Task) error {

	fmt.Println("Input task number to complete:")
	var num int
	fmt.Scanln(&num)
	if num > 0 && num <= len(*taskList) {
		(*taskList)[num-1].Status = true
	} else {
		err := errors.New("invalid input")
		log.Println(err)
		return err
	}
	log.Println("Статус задачи успешно изменён")
	return nil
}

func DeleteTask(taskList *[]Task) error {
	fmt.Println("Input task number to delete:")
	var num int
	fmt.Scan(&num)
	if num > len(*taskList) {
		err := errors.New("Input task number to delete is greater than the number of tasks")
		log.Println(err)
		return err
	}
	if num > 0 && num <= len(*taskList) {
		*taskList = append((*taskList)[:num-1], (*taskList)[num:]...)
		fmt.Println("Task deleted!")
	}
	log.Println("Задача успешно добавлена")
	return nil
}

func ViewTasks(taskList []Task) error {
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
