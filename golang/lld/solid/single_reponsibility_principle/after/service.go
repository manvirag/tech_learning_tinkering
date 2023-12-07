package after

import "fmt"

type IService interface {
	service()
}

func (c *App) service() {
	fmt.Println("this is the business logic")
}
