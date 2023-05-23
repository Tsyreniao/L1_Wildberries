package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Используем библиотеку math/big для работы с большими числами
	// Инициализируем переменные
	a := big.NewInt(26784390290532438)
	b := big.NewInt(1134890203808390)

	// Умножение больших чисел
	multiplyResult := new(big.Int)
	multiplyResult.Mul(a, b)
	fmt.Println("Умножение:", multiplyResult)

	// Деление больших чисел
	divideResult := new(big.Int)
	divideResult.Div(a, b)
	fmt.Println("Деление:", divideResult)

	// Сложение больших чисел
	addResult := new(big.Int)
	addResult.Add(a, b)
	fmt.Println("Сложение:", addResult)

	// Вычитание больших чисел
	subtractResult := new(big.Int)
	subtractResult.Sub(a, b)
	fmt.Println("Вычитание:", subtractResult)
}
