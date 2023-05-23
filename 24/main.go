package main

import (
	"fmt"
	"math"
)

// Структура точки
type Point struct {
	x int
	y int
}

// Функция нахождения расстояния м/у двумя точками
func distance(a, b Point) float64 {
	// Находим два катета ac и bc
	ac := a.x - b.x
	bc := a.y - b.y
	// Находим гипотенузу
	ab := math.Sqrt(float64(ac*ac + bc*bc))
	return ab
}

func main() {
	// Инициализируем точки
	var point1 = Point{
		x: -12,
		y: 12,
	}
	var point2 = Point{
		x: -13,
		y: -4,
	}

	// Выводим результат
	fmt.Println(distance(point1, point2))
}
