package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	ch := generator()
	for i := 0; i < 5; i++ {
		value := <-ch
		fmt.Println(value)
	}
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Println("Total Execution Time Of Program in microsecond: %+v", elapsedTime.Microseconds())
}

func generator() <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	return ch
}
