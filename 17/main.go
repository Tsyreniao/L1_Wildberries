package main

import (
	"fmt"
	"sort"
)

// Функция бинарного поиска позиции элемента
func binarySearch(array []int, num int, low, high int) int {
	var middle int = (high-low)/2 + low // Находим среднюю позицию
	var ans int
	// Находим область где находится элемент и запускаем поиск в этой области
	if num < array[middle] { // Элемент меньше находящегося посередине
		ans = binarySearch(array, num, low, middle)
	} else if num > array[middle] { // Элемент больше находящегося посередине
		ans = binarySearch(array, num, middle, high)
	} else { // Элемент находится посередине
		ans = middle
	}
	return ans
}

// Точка входа в поиск
func search(array []int, num int) int {
	// Задаются значения охватывающие весь массив
	return binarySearch(array, num, 0, len(array)-1)
}

func main() {
	// Создаем массив
	var array = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4}
	sort.Ints(array) // Сортируем

	// Выводим
	fmt.Println(array)

	// Бинарный поиск позиции элемента
	var pos int = search(array, 6)

	// Выводим результат
	fmt.Println(pos, array[pos])
}
