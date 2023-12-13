package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)

	go func() {
		time.Sleep(time.Second * 3)
		ch <- 4
	}()

	select {
	case val := <-ch:
		fmt.Println("channel value %+v", val)
	case other := <-time.After(time.Second * 1):
		fmt.Println("channel time out after one second %+v", other)
	}
}
