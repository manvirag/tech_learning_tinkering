package models

import (
	"fmt"
	"os"
	"sync"
)

type PrinterStratedy interface {
	Log(mesage string)
}

type ConsolePrinter struct{}

func (c *ConsolePrinter) Log(mesage string) {
	fmt.Println(mesage)
}

type FilePrinter struct{
	sync.Mutex
}

func (c *FilePrinter) Log(mesage string) {

	f, err := os.OpenFile("application.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()
	c.Lock()
	defer c.Unlock()
	if _, err = f.WriteString(mesage + "\n"); err != nil {
		fmt.Println("Error writing to file:", err)
	}
}

type DatabasePrinter struct {
}

func (c *DatabasePrinter) Log(mesage string) {
	// Can be implemented in the future
}
