package main

import "fmt"

// Функция переворота строки
func reverseString(s string) string {
	// Конвертируем строку в массив рун
	var runes = []rune(s)

	// Получаем длину массива
	var l = len(runes) - 1

	// Переворачиваем массив
	for i := 0; i < l/2+1; i++ {
		runes[i], runes[l-i] = runes[l-i], runes[i]
	}

	// Выводим как строку
	return string(runes)
}

func main() {
	var s string = "도시락 главрыба — абырвалг 도시락"

	// Вывод строки
	fmt.Println(s)

	// Переворот строки
	s = reverseString(s)

	// Вывод результата
	fmt.Println(s)
}
