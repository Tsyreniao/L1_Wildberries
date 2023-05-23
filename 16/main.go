package main

import "fmt"

// Функция для перестановки всех значений относительно значения pivot
func partition(array []int, low, high int) ([]int, int) {
	pivot := array[high] // Элемент с которым сравнивается область массива
	wall := low          // Стена - позиция относительно которой будут расположены элементы
	// Элементы меньше pivot будут левее стены, элементы больше pivot будут правее стены
	for j := low; j < high; j++ {
		if array[j] < pivot { // Если выбранный элемент меньше pivot, то
			// Элемент тановится на место wall
			array[wall], array[j] = array[j], array[wall]
			// wall передвигается дальше
			wall++
		}
	}
	array[wall], array[high] = array[high], array[wall] // Pivot и wall меняются местами
	return array, wall                                  // далее передается массив и позиция стены
}

// Функция сортировки области массива
func quickSort(array []int, low, high int) []int {
	if low < high { // Пока облать для сортировки существует, то сортируем
		var wall int
		array, wall = partition(array, low, high) // Получаем массив и позиция стены относительно которой отсортирован массив
		array = quickSort(array, low, wall-1)     // Сортировка левее стены
		array = quickSort(array, wall+1, high)    // Сортировка правее стены
	}
	return array
}

// Точка входа в сортировку
func sort(array []int) []int {
	// Задаются значения охватывающие весь массив
	return quickSort(array, 0, len(array)-1)
}

func main() {
	// Создаем массив
	var array = []int{6, 3, 1, 8, 4, 2, 0, 2, 2, 1, 0, 2, 1, 3, 1, 4, 6, 8, 5, 4}

	// Выводим
	fmt.Println(array)

	array = sort(array)

	// Проверяем
	fmt.Println(array)
}
