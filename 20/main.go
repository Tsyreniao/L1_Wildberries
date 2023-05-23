package main

import (
	"fmt"
	"strings"
)

// Функция переворота строки
func reverseWordsString(s string) string {
	// Нарезаем строку на слова
	var words = strings.Split(s, " ")
	var l = len(words) - 1 // Получаем кол-во слов

	// Переворачиваем порядок слов
	for i := 0; i < l/2+1; i++ {
		words[i], words[l-i] = words[l-i], words[i]
	}

	// Записываем слова в строку
	var result string
	for _, v := range words {
		result += v + " "
	}

	// Удаляем лишний пробел
	result = result[:len(result)-1]

	// Выводим результат
	return result
}

func main() {
	var s string = "1 도시락 2 snow 3 dog 4 sun 5 sun 6 dog 7 snow 8 도시락 9"

	// Вывод строки
	fmt.Println(s)

	// Переворот слов в строке
	s = reverseWordsString(s)

	// Вывод результата
	fmt.Println(s)
}
