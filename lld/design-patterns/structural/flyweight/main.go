// Package main is the entry point of the Word Processor application.
package main

import "fmt"

// Word represents a single word in the document.
type Word struct {
	Text string
}

// WordProcessor is the main component managing the document.
type WordProcessor struct {
	WordsMap map[string]*Word
}

// NewWordProcessor creates a new WordProcessor instance.
func NewWordProcessor() *WordProcessor {
	return &WordProcessor{
		WordsMap: make(map[string]*Word),
	}
}

// GetWord retrieves or creates a Word instance for the given text.
func (wp *WordProcessor) GetWord(text string) *Word {
	if word, ok := wp.WordsMap[text]; ok {
		return word
	}
	wp.WordsMap[text] = &Word{Text: text}
	return wp.WordsMap[text]
}

// Document represents the entire document composed of words.
type Document struct {
	Words []*Word
}

// AddWord adds a word to the document.
func (d *Document) AddWord(word *Word) {
	d.Words = append(d.Words, word)
}

// PrintDocument prints the content of the document.
func (d *Document) PrintDocument() {
	for _, word := range d.Words {
		fmt.Print(word.Text + " ")
	}
	fmt.Println()
}

func main() {
	wordProcessor := NewWordProcessor()

	document := Document{}
	document.AddWord(wordProcessor.GetWord("Hello"))
	document.AddWord(wordProcessor.GetWord("World"))
	document.AddWord(wordProcessor.GetWord("Hello"))

	document.PrintDocument()
}
