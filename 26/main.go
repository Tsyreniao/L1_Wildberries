package main

import (
	"fmt"
	"strings"
)

func worker(s string) bool {
	// Переводим строку в нижний регистр
	s = strings.ToLower(s)
	// Создаем map
	var charMap = make(map[rune]bool)

	for _, v := range s {
		if charMap[v] { // Если символ уже существует возвращаем false
			return false
		}
		// Заносим символ в map
		charMap[v] = true
	}

	// Если все символы уникальные возвращаем true
	return true
}

func main() {
	// Инициализируем строку
	var array string = "abcdEe"

	// Выводим результат
	fmt.Println(worker(array))
}
