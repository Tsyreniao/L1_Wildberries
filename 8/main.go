package main

import "fmt"

// Функция установки i-ого бита в b(1/0)
func worker(data int64, i int, b byte) int64 {
	// Создаем маску используя побитовый сдвиг
	var mask int64 = 1 << i

	if b == 1 {
		// Применяем логическое ИЛИ
		// 0 | 0 = 0
		// 1 | 0 = 1
		// 0 | 1 = 1
		// 1 | 1 = 1
		data = data | mask
	} else {
		// Применяем логическое И НЕ
		// 0 &^ 0 = 0
		// 1 &^ 0 = 1
		// 0 &^ 1 = 0
		// 1 &^ 1 = 0
		data = data &^ mask
	}
	return data
}

func main() {
	var data int64
	var i int
	var b byte

	data = 63 // Число
	i = 3     // Установим 3й бит (отсчет с 0)
	b = 0     // Установим 1

	var newData = worker(data, i, b)

	// Выводим данные
	fmt.Printf("%b - old\n", data)
	fmt.Printf("%b - new\n", newData)
}
