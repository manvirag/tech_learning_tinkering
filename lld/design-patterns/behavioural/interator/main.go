package main

import "fmt"

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type Aggregate interface {
	CreateIterator() Iterator
}

type SliceCollection struct {
	items []interface{}
}

func NewSliceCollection(items []interface{}) *SliceCollection {
	return &SliceCollection{items: items}
}

func (c *SliceCollection) CreateIterator() Iterator {
	return &SliceIterator{
		items: c.items,
		index: 0,
	}
}

type SliceIterator struct {
	items []interface{}
	index int
}

func (it *SliceIterator) HasNext() bool {
	return it.index < len(it.items)
}

func (it *SliceIterator) Next() interface{} {
	if !it.HasNext() {
		return nil
	}
	val := it.items[it.index]
	it.index++
	return val
}

func main() {
	items := []interface{}{"one", "two", "three"}
	collection := NewSliceCollection(items)
	iter := collection.CreateIterator()

	for iter.HasNext() {
		fmt.Println(iter.Next())
	}
}
