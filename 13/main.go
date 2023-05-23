package main

import "fmt"

func main() {
	// Создаем два числа
	var num1 int = 1
	var num2 int = 2

	// Выводим
	fmt.Println(num1, num2)

	// Меняем местами
	num1, num2 = num2, num1

	// Проверяем
	fmt.Println(num1, num2)
}
