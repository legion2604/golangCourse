package main

import (
	"fmt"
	"golangCourse/todo/internal/start"
	"golangCourse/todo/internal/task"
)

func main() {
	fmt.Println("Hello it's your ToDo app!")
	taskList := make([]task.Task, 0)
	start.Start(taskList)
}
