package main

import (
	"fmt"
)

// Command interface
type Command interface {
	Execute()
	Undo()
}

// Receiver - the object the command operates on
type Document struct {
	content string
}

func (d *Document) Write(text string) {
	d.content += text
}

func (d *Document) Erase() {
	if len(d.content) > 0 {
		d.content = d.content[:len(d.content)-1]
	}
}

func (d *Document) Print() {
	fmt.Println("Content:", d.content)
}

// Concrete Command - WriteCommand
type WriteCommand struct {
	document *Document
	text     string
}

func NewWriteCommand(document *Document, text string) *WriteCommand {
	return &WriteCommand{document, text}
}

func (wc *WriteCommand) Execute() {
	wc.document.Write(wc.text)
}

func (wc *WriteCommand) Undo() {
	wc.document.Erase()
}

// Concrete Command - EraseCommand
type EraseCommand struct {
	document *Document
}

func NewEraseCommand(document *Document) *EraseCommand {
	return &EraseCommand{document}
}

func (ec *EraseCommand) Execute() {
	ec.document.Erase()
}

func (ec *EraseCommand) Undo() {
	// For simplicity, no undo action for EraseCommand in this example
}

// Invoker
type Invoker struct {
	commands []Command
}

func (i *Invoker) ExecuteCommand(command Command) {
	command.Execute()
	i.commands = append(i.commands, command)
}

func (i *Invoker) Undo() {
	if len(i.commands) > 0 {
		lastCommand := i.commands[len(i.commands)-1]
		lastCommand.Undo()
		i.commands = i.commands[:len(i.commands)-1]
	}
}

func main() {
	// Client code

	document := &Document{}
	invoker := &Invoker{}

	// Writing commands
	writeCommand1 := NewWriteCommand(document, "Hello, ")
	writeCommand2 := NewWriteCommand(document, "World!")

	// Execute and store commands
	invoker.ExecuteCommand(writeCommand1)
	invoker.ExecuteCommand(writeCommand2)

	// Print document content
	document.Print() // Output: Content: Hello, World!

	// Undo the last command
	invoker.Undo()

	// Print document content after undo
	document.Print() // Output: Content: Hello,

	// Redo the undone command
	invoker.ExecuteCommand(writeCommand2)

	// Print document content after redo
	document.Print() // Output: Content: Hello, World!
}
