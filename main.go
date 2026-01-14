package main

import (
	"fmt"
	"golangCourse/gorutine/channel"
)

var a int

func main() {
	go channel.Problem1()
	for i := 1; i <= 5; i++ {
		fmt.Println(<-channel.ChProblem1)
	}

	channel.Problem2()
	for a := range channel.ChProblem2 {
		fmt.Println(a)
	}

}
