package main

import "fmt"

//--------------------------------------------------------------------------------
// Структура Square
type Square struct {
	Length float64
}

// Метод Square для нахождения площади
func (s *Square) CalculateArea() float64 {
	return s.Length * s.Length
}

//--------------------------------------------------------------------------------
// Структура Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Метод Rectangle для нахождения площади
func (r *Rectangle) CalculateArea() float64 {
	return r.Width * r.Height
}

//--------------------------------------------------------------------------------
// Интерфейс Shape
type Shape interface {
	CalculateArea() float64
}

// Адаптер ShapeAdapter для адаптации Square и Rectangle к интерфейсу Shape
type ShapeAdapter struct {
	Shape Shape
}

// Метод CalculateArea для ShapeAdapter
func (adapter *ShapeAdapter) CalculateArea() float64 {
	// Используем метод CalculateArea() у объектов удоволетворяющих интерфесу Shape
	return adapter.Shape.CalculateArea()
}

func main() {
	// Создаем квадрат и прямоугольник
	var square = Square{Length: 5}
	var rectangle = Rectangle{Width: 4, Height: 6}

	// Создаем адаптеры на основе квадрата и прямоугольника
	var squareAdapter = ShapeAdapter{Shape: &square}
	var rectangleAdapter = ShapeAdapter{Shape: &rectangle}

	// Вычисляем площади
	fmt.Println(squareAdapter.CalculateArea())
	fmt.Println(rectangleAdapter.CalculateArea())
}
