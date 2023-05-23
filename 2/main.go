package main

import (
	"fmt"
	"sync"
)

// Вычисление квадрата числа
func squareOfNumber(wg *sync.WaitGroup, number int) {
	// Сообщаем sync.WaitGroup что процесс завершится, перед выходом из функции
	defer wg.Done()

	// Вывод квадратов чисел в stdout
	fmt.Printf("%v^2 = %v\n", number, number*number)
}

func main() {
	// Создаем sync.WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup
	array := []int{2, 4, 6, 8, 10}

	for _, v := range array {
		// Добавляем 1 процесс в sync.WaitGroup и запускаем горутину
		wg.Add(1)
		go squareOfNumber(&wg, v)
	}

	// Ожидание завершения всех горутин
	wg.Wait()
}
