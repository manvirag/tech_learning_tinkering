package after

import "fmt"

type IDao interface {
	dao()
}

func (c *App) dao() {
	fmt.Println("this person will connect to db")
}
