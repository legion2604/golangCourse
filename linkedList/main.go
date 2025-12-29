package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := LinkedList{}
	arr.Add(12) // корект
	arr.Print()
	arr.Add(3) // корект
	arr.Print()
	arr.Add("24") // вызывет panic из-за разных типов
	arr.Print()
}

type LinkedList []Node

type Node struct {
	Value any
	Next  *Node
}

func (l *LinkedList) Add(value any) {
	newValue := Node{Value: value, Next: nil}
	if len(*l) == 0 {
		*l = append(*l, newValue)
		return
	}
	if reflect.TypeOf((*l)[len(*l)-1].Value) != reflect.TypeOf(newValue.Value) {
		panic("Неправельный тип элемента")
	}
	(*l)[len(*l)-1].Next = &newValue
	*l = append(*l, newValue)

}

func (l *LinkedList) Remove(index int) {
	newList := append((*l)[:index], (*l)[index+1:]...)
	*l = newList
}

func (l *LinkedList) Print() {
	for _, node := range *l {
		fmt.Print(node.Value, " ")
	}
	fmt.Println()
}
