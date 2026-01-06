package gorutine

import (
	"fmt"
	"math/rand"
	"time"
)

var dishes = []string{"coffee", "steak", "omelette", "pancakes", "water"}
var report []string

func order() map[int]string {
	res := make(map[int]string)
	for i := 0; i < rand.Intn(10); i++ {
		res[i] = dishes[rand.Intn(len(dishes))]
	}
	return res
}

func cookingProcess(id int, dish string) {
	cookingTime := rand.Intn(2000)
	time.Sleep(time.Duration(cookingTime) * time.Millisecond)
	report = append(report, fmt.Sprintf("Заказ номер #%d был на: %s", id, dish))
	fmt.Printf("Заказ номер #%d запущен.\n", id)
	fmt.Printf("Заказ номер #%d завершен за %d мс\n", id, cookingTime)
}

func Start() {
	orders := order() // заказы
	for k, v := range orders {
		go cookingProcess(k, v)
	}
	time.Sleep(time.Duration(len(orders)*2000) * time.Millisecond)
	fmt.Println(report) // список блюд на сегодня, отчёт
	fmt.Println(orders) // заказы с id
}
