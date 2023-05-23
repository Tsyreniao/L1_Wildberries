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

// Структура данных для потокобезопасного доступа к map
type MyMap struct {
	mx sync.RWMutex   // Mutex гарантирует, что только одна горутина будет иметь доступ к данным в момент времени
	m  map[string]int // map
}

// Функция для получения данных из map
func (myMap *MyMap) Load(key string) (int, bool) {
	myMap.mx.RLock()         // Блокировка доступа для чтения
	defer myMap.mx.RUnlock() // Разблокировка доступа для чтения
	val, ok := myMap.m[key]  // Получение данных
	return val, ok
}

// Функция для сохранения данных в map
func (myMap *MyMap) Store(key string, val int) {
	myMap.mx.Lock()         // Блокировка доступа для запсиси
	defer myMap.mx.Unlock() // Разблокировка доступа для запсиси
	myMap.m[key] = val      // Сохранение
}

func worker(ctx context.Context, wg *sync.WaitGroup, myMap *MyMap, i int) {
	defer wg.Done() // Сообщаем sync.WaitGroup что процесс завершится, перед выходом из функции
	var j int = 0
	for {
		select {
		case <-ctx.Done(): // Контекст завершен
			fmt.Println("Worker stopped by context")
			return
		default: // Кейс по умолчанию
			data := rg.Intn(10000)
			myMap.Store("worker."+fmt.Sprint(i)+"."+fmt.Sprint(j), data) // Сохранение данных в map
			fmt.Println("key: worker."+fmt.Sprint(i)+"."+fmt.Sprint(j)+" data: ", data)
			time.Sleep(5 * time.Millisecond)
			j++
		}
	}
}

func main() {
	// Создаем контекст с возможностью завершения
	ctx, cancel := context.WithCancel(context.Background())

	var myMap = MyMap{
		m: make(map[string]int),
	}

	// Создаем sync.WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ { // Всего запускается 5 горутин
		// Добавляем 1 процесс в sync.WaitGroup и запускаем горутину
		wg.Add(1)
		go worker(ctx, &wg, &myMap, i)
	}

	go func() {
		time.Sleep(30 * time.Millisecond)
		cancel() // Завершение контекста
	}()

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод данных
	fmt.Println("Data from map:")
	for key, val := range myMap.m {
		fmt.Printf("key: %v data: %v\n", key, val)
	}

	fmt.Println("Main stopped")
}
