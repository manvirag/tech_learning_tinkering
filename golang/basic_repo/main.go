package main

import (
	"fmt"

	"github.com/google/uuid"
	"rsc.io/quote"
)

func main() {

	fmt.Println(quote.Go())

	for i := 0; i < 5; i++ {
		println(uuid.New().String())
	}
}
