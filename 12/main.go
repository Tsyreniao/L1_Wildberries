package main

import "fmt"

func main() {

	// Создаем массив
	var array = []string{"cat", "cat", "dog", "cat", "tree"}

	// Множество уникальных элементов
	var hash = make(map[string]bool)

	// Проходим по массиву и сохраняем значения в Map
	for _, v := range array {
		hash[v] = true
	}

	// Вывод множества
	for k := range hash {
		fmt.Println(k)
	}
}
