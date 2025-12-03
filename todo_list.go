package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

type task struct {
	description string
	status      bool
	createdAt   string
}

func main() {
	fmt.Println("Hello it's your ToDo app!")
	taskList := make([]task, 0)

	for {
		fmt.Println("\n", "1. Add Task", "\n", "2. View Tasks", "\n", "3. Complete the task", "\n", "4. Delete Task", "\n", "5. Exit")
		var input int
		fmt.Scan(&input)

		if input == 1 {
			fmt.Println("Input task description:")
			desc := Print()
			date := "1999-12-31"
			time.Parse(layoutISO, date)
			newTask := task{description: desc, status: false, createdAt: date}
			taskList = append(taskList, newTask)
			fmt.Println("Task added!")
		} else if input == 2 {
			fmt.Println("Your Tasks:")
			for i, t := range taskList {
				fmt.Println(i+1, " Description:", t.description, "| Status: ", t.status, " | Created At: ", t.createdAt)
			}
		} else if input == 3 {
			fmt.Println("Input task number to complete:")
			var num int
			fmt.Scanln(&num)
			if num > 0 && num <= len(taskList) {
				taskList[num-1].status = true
				fmt.Println("Task marked as completed!")
			} else {
				fmt.Println("Task not complete! Enter valid task number.")
			}
		} else if input == 4 {
			fmt.Println("Input task number to delete:")
			var num int
			fmt.Scanln(&num)
			if num > 0 && num <= len(taskList) {
				taskList = append(taskList[:num-1], taskList[num:]...)
				fmt.Println("Task deleted!")
			}
		} else if input == 5 {
			fmt.Println("Bye!")
			break
		} else {
			fmt.Println("Invalid input")
		}
	}
}

func Print() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = text[:len(text)-1]

	return text
}
