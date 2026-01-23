package _map

import "fmt"

func Start() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2, 3, 7}
	res := make(map[int]int)
	res[1] = 1
	for _, v := range arr {
		res[v] = res[v] + 1
	}
	for k, v := range res {
		if v != 1 {
			fmt.Println(k, false)
		} else {
			fmt.Println(k, true)
		}
	}
}
