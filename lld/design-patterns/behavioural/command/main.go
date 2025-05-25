package main

import "fmt"

// Command interface
type Command interface {
	Execute()
}

// Receiver
type Light struct{}

func (l *Light) On() {
	fmt.Println("The light is ON")
}

func (l *Light) Off() {
	fmt.Println("The light is OFF")
}

// Concrete Command to turn the light ON
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

// Concrete Command to turn the light OFF
type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

// Invoker
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(c Command) {
	r.command = c
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

// Client
func main() {
	light := &Light{}

	lightOn := &LightOnCommand{light}
	lightOff := &LightOffCommand{light}

	remote := &RemoteControl{}

	remote.SetCommand(lightOn)
	remote.PressButton() // Output: The light is ON

	remote.SetCommand(lightOff)
	remote.PressButton() // Output: The light is OFF
}
