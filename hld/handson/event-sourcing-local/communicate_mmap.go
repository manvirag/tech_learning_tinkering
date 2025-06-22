package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"syscall"
	"time"
)

// Event represents an event
type Event struct {
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}

// Constants
const (
	logFilePath = "eventlog.dat"
	fileSize    = 1024 * 1024 // 1 MB
)

// Write event to mmap-ed file
func writeEvent(mmapData []byte, offset *int64, mu *sync.Mutex, e Event) error {
	mu.Lock()
	defer mu.Unlock()

	// Serialize event
	eventBytes, err := json.Marshal(e)
	if err != nil {
		return err
	}

	// Write length prefix
	binary.LittleEndian.PutUint32(mmapData[*offset:], uint32(len(eventBytes)))
	*offset += 4

	// Write event data
	copy(mmapData[*offset:], eventBytes)
	*offset += int64(len(eventBytes))

	return nil
}

// Read events from mmap-ed file
func readEvents(mmapData []byte) {
	var offset int64 = 0

	for {
		if offset+4 > int64(len(mmapData)) {
			break
		}

		// Read length prefix
		length := binary.LittleEndian.Uint32(mmapData[offset:])
		if length == 0 {
			break
		}
		offset += 4

		if offset+int64(length) > int64(len(mmapData)) {
			break
		}

		// Read event
		data := mmapData[offset : offset+int64(length)]
		var e Event
		err := json.Unmarshal(data, &e)
		if err != nil {
			fmt.Println("Error reading event:", err)
			break
		}
		fmt.Printf("Read event: %+v\n", e)
		offset += int64(length)
	}
}

func main() {
	// Create or open file
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Ensure file size
	err = file.Truncate(fileSize)
	if err != nil {
		panic(err)
	}

	// Memory map the file
	mmapData, err := syscall.Mmap(int(file.Fd()), 0, fileSize, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED)
	if err != nil {
		panic(err)
	}
	defer syscall.Munmap(mmapData)

	var offset int64 = 0
	var mu sync.Mutex

	// Writer goroutine (simulating event producer)
	go func() {
		for i := 0; i < 5; i++ {
			e := Event{
				Timestamp: time.Now().Unix(),
				Message:   fmt.Sprintf("Event %d", i),
			}
			err := writeEvent(mmapData, &offset, &mu, e)
			if err != nil {
				fmt.Println("Write error:", err)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	// Reader goroutine (simulating event consumer)
	go func() {
		for {
			mu.Lock()
			readEvents(mmapData)
			mu.Unlock()
			time.Sleep(2 * time.Second)
		}
	}()

	// Run long enough to see output
	time.Sleep(12 * time.Second)
}


/*

💡 What is mmap?
mmap stands for memory map (or memory-mapped file).
It is a mechanism provided by the operating system that maps a file (or a device or anonymous memory) directly into a process's address space.

👉 This means:

The contents of a file (or region of memory) can be accessed as if it were just an array or pointer in your program.

You can read or write to the file by simply reading or writing memory.

📝 How does mmap work?
🔹 Instead of using traditional I/O (e.g., read, write), the OS:

Maps the file into virtual memory.

Any changes you make to that memory can be reflected in the file (depending on flags).

🔹 Internally, the OS uses page tables to link file contents to memory pages.


for reliability, we can forcely write at time of appending, while reading via mmap 
*/
