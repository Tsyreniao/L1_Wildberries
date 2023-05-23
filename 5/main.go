package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Создание *rand.Rand для генерации рандомных значений
var rg = rand.New(rand.NewSource(time.Now().Unix()))

func writer(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done() // Сообщаем sync.WaitGroup что процесс завершится, перед выходом из функции
	defer close(ch) // Закрытие канала для передачи данных, перед выходом из функции
	for {
		data := rg.Intn(100000) // Генерация рандомного числа
		select {
		case <-ctx.Done(): // Контекст завершен
			fmt.Println("Writer stopped")
			return
		case ch <- data: // Запись в канал
			fmt.Println("Data sent: ", data)
		}
	}
}

func worker(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done() // Сообщаем sync.WaitGroup что процесс завершится, перед выходом из функции
	for {
		select {
		case <-ctx.Done(): // Контекст завершен
			fmt.Println("Worker stopped")
			return
		default: // Кейс по умолчанию
			data, open := <-ch
			if open {
				// Если open == true (канал открыт), то выводим данные
				fmt.Println("Data received:", data)
			}
			// Если open == false (канал закрыт), то идем на следующую итерацию
		}
	}
}

func main() {
	// Считываем кол-во секунд
	var n int
	fmt.Scan(&n)

	// Создаем контекст с возможностью завершения
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(n))
	defer cancel()

	// Создаем канал для передачи данных
	ch := make(chan int)

	// Создаем sync.WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup
	// Добавляем 1 процесс в sync.WaitGroup и запускаем горутину
	wg.Add(1)
	go writer(ctx, &wg, ch)

	// Добавляем 1 процесс в sync.WaitGroup и запускаем горутину
	wg.Add(1)
	go worker(ctx, &wg, ch)

	// Ожидание завершения всех горутин
	wg.Wait()
	fmt.Println("Main stopped")
}
