package main

import "fmt"

// Структура Human
type Human struct {
	name string
	age  uint8
}

// Метод Human
func (h *Human) Hello() {
	fmt.Printf("Hi, my name is %v\n", h.name)
}

// Встраивание структуры Human в Action
type Action struct {
	// Анонимное поле без имени, которое позволяет Action использовать методы Human
	Human
}

// Метод Action
func (a *Action) Run() {
	fmt.Printf("%v is running\n", a.name)
}

func main() {
	// Инициализация объекта Human
	human := Human{
		name: "Tsyren",
		age:  22}

	// Вызов метода Human
	human.Hello()

	// Инициализация объекта Action
	runner := Action{
		Human: human,
	}

	// Вызов метода Human
	runner.Hello()

	// Вызов метода Action
	runner.Run()
}
