package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Создание *rand.Rand для генерации рандомных значений
var rg = rand.New(rand.NewSource(time.Now().Unix()))

func writer(ctx context.Context, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done() // Сообщаем sync.WaitGroup что процесс завершится, перед выходом из функции
	defer close(ch) // Закрытие канала для передачи данных
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

func worker(ctx context.Context, wg *sync.WaitGroup, ch chan int, i int) {
	defer wg.Done() // Сообщаем sync.WaitGroup что процесс завершится, перед выходом из функции
	for {
		select {
		case <-ctx.Done(): // Контекст завершен
			fmt.Printf("Worker №%v stopped\n", i)
			return
		default: // Кейс по умолчанию
			data, open := <-ch
			if open {
				// Если open == true (канал открыт), то выводим данные
				fmt.Printf("Worker №%v\tData: %v\n", i, data)
			}
			// Если open == false (канал закрыт), то идем на следующую итерацию
		}
	}
}

func main() {
	// Создаем контекст с возможностью завершения
	ctx, cancel := context.WithCancel(context.Background())
	// Создаем канал, который получает сигналы из ОС
	osSigCh := make(chan os.Signal)
	signal.Notify(osSigCh, syscall.SIGINT) // syscall.SIGINT - Сигнал прерывания пользователем(Ctrl+C)
	go func() {
		// Ожидаем пока в канал не придут данные
		<-osSigCh
		// Завершение контекста
		cancel()
	}()

	// Считываем кол-во воркеров
	var n int
	fmt.Scan(&n)

	// Создаем канал для передачи данных
	ch := make(chan int)

	// Создаем sync.WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup
	// Добавляем 1 процесс в sync.WaitGroup и запускаем горутину
	wg.Add(1)
	go writer(ctx, &wg, ch)

	for i := 0; i < n; i++ {
		// Добавляем 1 процесс в sync.WaitGroup и запускаем горутину
		wg.Add(1)
		go worker(ctx, &wg, ch, i)
	}

	// Ожидание завершения всех горутин
	wg.Wait()
	fmt.Println("Main stopped")
}
