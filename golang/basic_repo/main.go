package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/manvirag982/tech_learning_tinkering/golang/publish_repo/package_one"
	"github.com/manvirag982/tech_learning_tinkering/golang/publish_repo/package_two"
	"rsc.io/quote"
)

func main() {

	fmt.Println(quote.Go())
	package_one.HelloPackageOne()
	package_two.HelloPackageTwo()

	for i := 0; i < 5; i++ {
		println(uuid.New().String())
	}
}
