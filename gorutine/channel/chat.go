package channel

import "fmt"

type massage struct {
	massage string
	userId  int
}

var userChan = make(chan massage)

func Start() {
	arr1 := []massage{{massage: "hello", userId: 1}, {massage: "how are you?", userId: 1}, {massage: "I`m fine too, thanks", userId: 1}, {massage: "what are you doing?", userId: 1}}
	arr2 := []massage{{massage: "hello", userId: 2}, {massage: "I, fine thanks, and what about you?", userId: 2}, {massage: "nothing 😊", userId: 2}}
	go func() {
		user(arr1)
		user(arr2)
		close(userChan)
	}()

	for {
		printAllMassage()
	}
}

func printAllMassage() {
	for {
		mass, ok := <-userChan
		if !ok {
			break
		}
		fmt.Println("User:", mass.userId, "Massage:", mass.massage)
	}
}

func user(arr []massage) {
	for _, v := range arr {
		userChan <- v
	}
}

/*
а это сам сделал!!!
*/
