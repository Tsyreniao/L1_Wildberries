package main

import (
	"fmt"
	"sync"
)

// Вычисление квадрата числа
func squareOfNumber(wg *sync.WaitGroup, ch chan int, number int) {
	// Сообщаем sync.WaitGroup что процесс завершится, перед выходом из функции
	defer wg.Done()
	// Записываем ответ в канал
	ch <- number * number
}

func main() {
	// Создаем sync.WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup
	// Создание канала, куда будут записываться квадраты чисел
	ch := make(chan int)
	array := []int{2, 4, 6, 8, 10}

	for _, v := range array {
		// Добавляем 1 процесс в sync.WaitGroup и запускаем горутину
		wg.Add(1)
		go squareOfNumber(&wg, ch, v)
	}

	go func() {
		// Ожидание завершения всех горутин
		wg.Wait()
		// Закрытие канала происходит после завершения всех горутин
		close(ch)
	}()

	var sum int = 0
	// Тут программа ждет пока закроют канал
	for v := range ch {
		// Канал закрыт, читаем данные с канала
		sum += v
	}
	// Тот же цикл, но в более понятном виде
	// for {
	// 	v, open := <-ch //если канал закрыли, то open = false
	// 	if !open {
	// 		break
	// 	}
	// 	sum += v
	// }

	fmt.Println(sum)
}
