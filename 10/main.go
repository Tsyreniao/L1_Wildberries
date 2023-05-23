package main

import "fmt"

func main() {
	// Создаем массив данных
	var array = []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Создаем map где будем хранить отсортированные данные
	var group = make(map[int][]float32)

	for _, v := range array {
		round := int(v/10) * 10                // Получаем округленное значение
		group[round] = append(group[round], v) // Сохраняем данные в map в соответствующий массив с соответствующим ключом
	}

	// Выводим данные
	for k, v := range group {
		fmt.Printf("%v:%v\n", k, v)
	}
}
