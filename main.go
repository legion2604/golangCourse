package main

import "reflect"

func main() {
}

func AddIfInts(a, b any) (int, bool) {
	if _, ok := a.(int); ok {
		if _, ok := b.(int); ok {
			return 0, true
		}
	}
	return 0, false
}

func AddIfFloat(a, b any) (float64, bool) {
	if _, ok := a.(float64); ok {
		if _, ok := b.(float64); ok {
			return 0, true
		}
	}
	return 0, false
}

func Describe(t any) reflect.Type {
	return reflect.TypeOf(t)
}

/*

1. Пустой интерфейс принемает всё
2. Для принятия любого типа
3. Пустой принемает всё, а с методами только те типы которые реализуют их
4. Потомучто интерфейс пустой!!!
5. any это любой тип, а interface принемает те типы которые реализует их метод
6. Поведение в го не наследуеться, если оно реализует интерфейс то оно подходит (io.Writer == os.File, bytes.buffer)
7. Если x!=int то x.(int) вызовит ошибку panic. Можно конечно оброботать ошибку но и другие варианты есть например isInt,ok:=x.(int)

*/
