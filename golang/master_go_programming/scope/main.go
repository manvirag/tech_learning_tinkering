package main

import (
	"fmt"
	"master_go_programming/scope"
)

// var glo int = 7 // package variable (we can access it in other file if we put package main there see integer.go file)
func main() {
	var x int = 5 // block variable
	fmt.Println(x)
	scope.Solve()
	// nn = "anubhav"
}
