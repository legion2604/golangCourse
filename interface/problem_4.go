package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := FilterInts([]any{"12", 3, 12.6})
	sum := Sum(arr)
	fmt.Println(arr, sum)
}

func FilterInts(xs []any) []int {
	var res []int

	for _, v := range xs {
		switch val := v.(type) {
		case int:
			res = append(res, val)
		case float64:
			res = append(res, int(val))
		case string:
			strToInt, ok := strconv.Atoi(val)
			if ok == nil {
				res = append(res, strToInt)
			} else {
				fmt.Println("string is not int, Idiot")
			}
		case bool:
			fmt.Println("its bool not number")
		}
	}

	return res
}

func Sum(xs []int) int {
	res := 0
	for _, x := range xs {
		res += x
	}
	return res
}
