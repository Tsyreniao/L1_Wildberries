package main

import (
	"fmt"
	"reflect"
)

// определения типа используя рефлексию
func coolerDefineType(variable interface{}) {
	// Получаем информацию о типа переменной и выводим
	fmt.Printf("Переменная: %v - Тип: %v\n", variable, reflect.TypeOf(variable))
}

// определения типа используя утверждение типа
func defineType(variable interface{}) {
	switch variable.(type) { // Используя switch сравниваем тип переменной конкретными типами
	case int:
		fmt.Printf("Переменная: %v - тип: int\n", variable)
	case string:
		fmt.Printf("Переменная: %v - тип: string\n", variable)
	case bool:
		fmt.Printf("Переменная: %v - тип: bool\n", variable)
	case chan interface{}:
		fmt.Printf("Переменная: %v - тип: chan interface{}\n", variable)
	default:
		fmt.Printf("Переменная: %v - тип: неизвестен\n", variable)
	}
}

func main() {
	// Создаем переменную типа interface{}
	var variable interface{}

	//Задаем значение и вызываем функцию определения типа
	variable = 22
	coolerDefineType(variable)

	//Задаем значение и вызываем функцию определения типа
	variable = "Tsyren"
	coolerDefineType(variable)

	//Задаем значение и вызываем функцию определения типа
	variable = true
	coolerDefineType(variable)

	//Задаем значение и вызываем функцию определения типа
	ch := make(chan interface{})
	variable = ch
	coolerDefineType(variable)
}
