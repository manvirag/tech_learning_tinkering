package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/exp/mmap"
)

const (
	eventFile = "events.dat"
)

// CreateEvent creates a new event and appends it to the memory-mapped file
func CreateEvent(id int32, name string) error {
	f, err := os.OpenFile(eventFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	event := fmt.Sprintf("%d:%s\n", id, name)
	if _, err := f.WriteString(event); err != nil {
		return err
	}
	return nil
}

// UpdateEvent updates an existing event in the memory-mapped file
func UpdateEvent(id int32, newName string) error {
	f, err := os.OpenFile(eventFile, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	r, err := mmap.Open(eventFile)
	if err != nil {
		return err
	}
	defer r.Close()

	// Read the entire file into memory
	p := make([]byte, r.Len())
	if _, err := r.ReadAt(p, 0); err != nil {
		return err
	}

	lines := strings.Split(string(p), "\n")
	newLines := ""
	for _, line := range lines {
		var currentID int32
		var currentName string
		n, _ := fmt.Sscanf(line, "%d:%s", &currentID, &currentName)

		if n == 2 && currentID == id {
			newLines += fmt.Sprintf("%d:%s\n", id, newName) // Update name
		} else if n == 2 {
			newLines += line + "\n" // Keep original line
		}
	}

	f.Truncate(0) // Clear the file
	f.Seek(0, 0)  // Reset the pointer to the beginning of the file
	if _, err := f.WriteString(newLines); err != nil {
		return err
	}
	return nil
}

// DeleteEvent deletes an event from the memory-mapped file
func DeleteEvent(id int32) error {
	f, err := os.OpenFile(eventFile, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	r, err := mmap.Open(eventFile)
	if err != nil {
		return err
	}
	defer r.Close()

	p := make([]byte, r.Len())
	if _, err := r.ReadAt(p, 0); err != nil {
		return err
	}

	lines := strings.Split(string(p), "\n")
	newLines := ""
	for _, line := range lines {
		var currentID int32
		fmt.Sscanf(line, "%d:", &currentID)

		if currentID != id { // Only keep lines that don't match the ID to delete
			newLines += line + "\n"
		}
	}

	f.Truncate(0) // Clear the file
	f.Seek(0, 0)  // Reset the pointer to the beginning of the file
	if _, err := f.WriteString(newLines); err != nil {
		return err
	}
	return nil
}

// ListEvents prints all events in the memory-mapped file
func ListEvents() error {
	r, err := mmap.Open(eventFile)
	if err != nil {
		return err
	}
	defer r.Close()

	p := make([]byte, r.Len())
	if _, err := r.ReadAt(p, 0); err != nil {
		return err
	}
	fmt.Println("Events:")
	fmt.Println(string(p))
	return nil
}

func main() {
	err := CreateEvent(1, "Ride Requested")
	if err != nil {
		log.Fatalf("Error creating event: %v", err)
	}

	err = CreateEvent(2, "Ride Accepted")
	if err != nil {
		log.Fatalf("Error creating event: %v", err)
	}

	err = ListEvents()
	if err != nil {
		log.Fatalf("Error listing events: %v", err)
	}

	err = UpdateEvent(1, "Ride Canceled")
	if err != nil {
		log.Fatalf("Error updating event: %v", err)
	}

	err = ListEvents()
	if err != nil {
		log.Fatalf("Error listing events: %v", err)
	}

	err = DeleteEvent(2)
	if err != nil {
		log.Fatalf("Error deleting event: %v", err)
	}

	err = ListEvents()
	if err != nil {
		log.Fatalf("Error listing events: %v", err)
	}
}
