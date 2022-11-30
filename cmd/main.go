package main

import (
	"fmt"
	"modul_20.2.1/pkg/pipeline"
	"time"
)

// Интервал очистки кольцевого буфера
const bufferDrainInterval = 30 * time.Second

// Размер кольцевого буфера
const bufferSize int = 10

func main() {

	fmt.Println("Для заполнения буфера введите целые числа")

	input := make(chan int)
	done := make(chan bool)

	go pipeline.Read(input, done)

	negativeFiltrCh := make(chan int)
	go pipeline.NegativeFiltrStageInt(input, negativeFiltrCh, done)

	notDivadedThreeCh := make(chan int)
	go pipeline.NotDivadedThreeFunc(negativeFiltrCh, notDivadedThreeCh, done)

	bufferedIntCh := make(chan int)

	go pipeline.BufferStageFunc(notDivadedThreeCh, bufferedIntCh, done, bufferSize, bufferDrainInterval)

	for {
		select {

		case data := <-bufferedIntCh:
			fmt.Println("Получены данные ... ", data)
		case <-done:
			return
		}
	}

}
