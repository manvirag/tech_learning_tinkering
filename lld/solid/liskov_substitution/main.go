// Example

package main

import "fmt"

// Liskov ...
type Liskov interface {
	LiskovFunc()
}

// Parent ...
type Parent struct{}

// Child ...   // this is not inheritance this is struct embedding -> https://www.tutorialspoint.com/composition-in-golang
// now child can access parent methods as well as field direct.
type Child struct {
	Parent
}

// LiskovFunc ...
func (p *Parent) LiskovFunc() {
	fmt.Println("It works")
}

// LiskovSubstitution ...
func LiskovSubstitution(lis Liskov) {
	lis.LiskovFunc()
}

func main() {
	ch := &Child{}
	par := &Parent{}
	LiskovSubstitution(ch)
	LiskovSubstitution(par)
}
