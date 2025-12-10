package main

import (
	"errors"
	"fmt"
)

func main() {
	a := 10
	a, err := plus10Limit(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(a)
}

func plus10Limit(a int) (int, error) {
	if a <= 0 {
		err := errors.New("a<=0")
		return 0, err
	}
	a += 10
	return a, nil
}
