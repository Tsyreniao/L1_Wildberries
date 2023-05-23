package main

import (
	"fmt"
	"sync"
)

func writer(wg *sync.WaitGroup, num chan<- int, ans chan<- int, value int) {
	defer wg.Done() // Сообщаем sync.WaitGroup что процесс завершится, перед выходом из функции
	// Записываем в каналы число и число * 2
	num <- value
	ans <- value * 2
}

func reader(wg *sync.WaitGroup, num <-chan int, ans <-chan int) {
	defer wg.Done() // Сообщаем sync.WaitGroup что процесс завершится, перед выходом из функции
	// Выводим ответ в stdout читая из каналов num и ans
	fmt.Printf("%v * 2 = %v\n", <-num, <-ans)
}

func main() {
	var array = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Создаем каналы
	num := make(chan int)
	ans := make(chan int)

	// Создаем sync.WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	for _, v := range array {
		// Добавляем 1 процесс в sync.WaitGroup и запускаем горутину
		wg.Add(1)
		go writer(&wg, num, ans, v)
		// Добавляем 1 процесс в sync.WaitGroup и запускаем горутину
		wg.Add(1)
		go reader(&wg, num, ans)
	}

	// Ожидание завершения всех горутин
	wg.Wait()
	close(num)
	close(ans)
}
