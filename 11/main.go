package main

import "fmt"

func main() {
	// Создаем массивы
	var array1 = []int{1, 1, 3, 3, 5, 5, 7, 7, 9, 9, 9, 12, 13, 14, 14, 20}
	var array2 = []int{1, 1, 9, 3, -1, 4, 20, 12, 12, 12}
	// Массив для пересечений
	var ans []int

	// Map для записи кол-ва повторений данных массива array1
	var hash = make(map[int]int)

	// Проходим по массиву array1 и записываем кол-во повторений значений в map
	for _, v := range array1 {
		hash[v] = +1
	}

	// Проходим по массиву array1 и проверяем есть-ли такиеже значения в map
	for _, v := range array2 {
		if hash[v] > 0 { // Если есть, то
			ans = append(ans, v) // Сохраняем значение
			hash[v] = -1         // Вычитаем повторение
		}
	}

	// Выводим массив с пересечениями
	for _, v := range ans {
		fmt.Printf("%v\n", v)
	}
}
