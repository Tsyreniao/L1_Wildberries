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

func worker(ctx context.Context, wg *sync.WaitGroup, ch chan bool, flag *bool) {
	defer wg.Done() // Сообщаем sync.WaitGroup что процесс завершится, перед выходом из функции
	//defer close(ch)
	for {
		select {
		case <-ctx.Done(): // Контекст завершен
			fmt.Println("Worker stopped by context")
			return
		case <-ch: // Получен сигнал остановки
			fmt.Println("Worker stopped by chanel")
			return
		default: // Кейс по умолчанию
			if *flag { // Флаг остановки == true
				fmt.Println("Worker stopped by flag")
				return
			}
			fmt.Println("Random data: ", rg.Intn(10000))
		}
	}
}

func main() {
	// Создаем контекст с возможностью завершения
	ctx, cancel := context.WithCancel(context.Background())

	// Создаем канал для сигнала остановки
	ch := make(chan bool)

	// Создание флага для остановки
	flag := false

	// Создаем sync.WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup
	// Добавляем 1 процесс в sync.WaitGroup и запускаем горутину
	wg.Add(1)
	go worker(ctx, &wg, ch, &flag)

	time.Sleep(30 * time.Millisecond)
	switch rnd := rg.Intn(3); rnd {
	case 0:
		cancel() // Завершение контекста
	case 1:
		ch <- true // Подаем в канал сигнал остановки
	case 2:
		flag = true // Поднят флаг остановки
	}

	// Ожидание завершения всех горутин
	wg.Wait()
	fmt.Println("Main stopped")
}
