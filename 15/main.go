package main

import "fmt"

var justString string

// Создаю большую строку
func createHugeString(size int) string {
	var hugeString string
	for i := 0; i < size; i++ {
		hugeString = hugeString + "F"
	}
	return hugeString
}

func someFunc() { // Создается большая строка, но используется только малая её часть
	// Можно изначально создавать не большую строку
	//v := createHugeString(100)
	v := createHugeString(1 << 10)

	// Возможно большая строка окажется не такой большой
	if len(v) < 100 { //тогда нужно сохранить всю строку
		justString = v
	} else {
		justString = v[:100]
	}
}

func main() {
	someFunc()
	fmt.Println(justString)
}
