package main

import (
	"fmt"
	"sync"
)

// Структура счетчик
type Counter struct {
	mx    sync.RWMutex // Мьютекс для блокировки и разблокировки доступа в конкурентной среде
	value int          // Значение счетчика
}

// Инкрементирование счетчика
func (c *Counter) Inc() {
	c.mx.Lock()         // Блокируем запись
	defer c.mx.Unlock() // Разблокируем запись

	c.value++
}

// Получение значения счетчика
func (c *Counter) Return() int {
	c.mx.RLock()         // Блокируем чтение
	defer c.mx.RUnlock() // Разблокируем чтение

	return c.value
}

func main() {
	// Инициализация счетчика
	var counter = Counter{
		value: 0,
	}

	// Создаем sync.WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	for i := 0; i < 100000; i++ {
		// Добавляем 1 процесс в sync.WaitGroup и запускаем горутину
		wg.Add(1)
		go func() {
			defer wg.Done() // Сообщаем sync.WaitGroup что процесс завершится, перед выходом из функции
			counter.Inc()
		}()
	}

	// Ожидание завершения всех горутин
	wg.Wait()

	// Вывод результата
	fmt.Println(counter.Return())

}
