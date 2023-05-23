package main

import (
	"fmt"
	"time"
)

// Функция sleep получает на вход кол-во секунд
func sleep(seconds int) {
	// Получаем значение из канала возвращаемого time.After, но никуда не записываем
	// Т.к. нас интересует только задержка времени
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	// Приостанавливаем выполнение на 2 секунды
	sleep(2)
	fmt.Println("Прошло 2 секунды")

	// Приостанавливаем выполнение на 5 секунд
	sleep(5)
	fmt.Println("Прошло 5 секунд")
}
