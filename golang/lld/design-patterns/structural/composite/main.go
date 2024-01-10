package main

import (
	"fmt"
)

// Component interface
type FileSystemComponent interface {
	List()
}

// Leaf: File
type File struct {
	Name string
}

func (f *File) List() {
	fmt.Println("File:", f.Name)
}

// Composite: Directory
type Directory struct {
	Name       string
	components []FileSystemComponent
}

func (d *Directory) List() {
	fmt.Println("Directory:", d.Name)
	for _, component := range d.components {
		component.List()
	}
}

func (d *Directory) Add(component FileSystemComponent) {
	d.components = append(d.components, component)
}

func main() {
	// Creating files
	file1 := &File{Name: "file1.txt"}
	file2 := &File{Name: "file2.txt"}

	// Creating directories
	dir1 := &Directory{Name: "dir1"}
	dir2 := &Directory{Name: "dir2"}

	// Adding files to directories
	dir1.Add(file1)
	dir2.Add(file2)

	// Creating a composite directory
	compositeDir := &Directory{Name: "compositeDir"}
	compositeDir.Add(dir1)
	compositeDir.Add(dir2)

	// Listing the composite directory
	compositeDir.List()

	// Output:
	// Directory: compositeDir
	// Directory: dir1
	// File: file1.txt
	// Directory: dir2
	// File: file2.txt
}
