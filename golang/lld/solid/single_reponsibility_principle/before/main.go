package before

import "fmt"

type IApp interface {
	service()
	dao()
}

type App struct {
	name string
}

func (c *App) service() {
	fmt.Println("this is the business logic")
}

func (c *App) dao() {
	fmt.Println("this person will connect to db")
}
