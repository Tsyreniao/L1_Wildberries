package main

import "fmt"

// Удаление i-ого элемента в слайсе
func deleteElement(s []int, i int) []int {
	copy(s[i-1:], s[i:]) // Сдвигаем правую часть на 1 позицию в лево
	s[len(s)-1] = 0      // Записываем нулевое значение в последний элемент
	s = s[:len(s)-1]     // Удаляем последний элемент
	return s
}

func main() {
	// Инициализация слайса
	var s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(s)

	s = deleteElement(s, 5)

	fmt.Println(s)
}
