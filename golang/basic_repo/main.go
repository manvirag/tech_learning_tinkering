package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/manvirag982/tech_learning_tinkering/golang/publish_repo"
	"rsc.io/quote"
)

func main() {

	fmt.Println(quote.Go())
	publish_repo.HelloWorld()

	for i := 0; i < 5; i++ {
		println(uuid.New().String())
	}
}
