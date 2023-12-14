We can use any primitive or user created type for sending events in channel(even functions).


#### Unbuffered Channels
- Can lead to program slowdown.
- Must wait for each message to be processed before sending another.

#### Buffered Channels
- Limited capacity.
- Maintains a message queue for goroutine consumption.
- Performance advantages:
    - Buffered channel empty: Receiving is blocked until a message is sent.
    - Buffered full: Sending is blocked until a message is received.
    - Partial: Both scenarios are possible.

### Closing Channels and Channel Behavior


- **Unbuffered Closed Channel:**
When you close an unbuffered channel, it means that no more values will be sent on the channel. If you try to send a value after closing, it will cause a panic. However, for receiving, the behavior is different:
    - Yields an immediate zero value of the element type.
    ```
      package main
      import (
      "fmt"
      )
    
      func main() {
      ch := make(chan int)
    
          close(ch) // Closing an unbuffered channel
    
          // Trying to send on a closed channel will cause a panic
          // go func() {
          // 	ch <- 42
          // }()
    
          // Receiving from a closed unbuffered channel
          val, ok := <-ch
          fmt.Println(val, ok) // Outputs: 0 false
        }
    ``` 
  
    

- **Buffered Closed Channel:**
  When you close a buffered channel, it again means that no more values will be sent on the channel. Sending on a closed buffered channel will also cause a panic. However, for receiving:

- It yields all values that are currently in the channel's queue.
After exhausting these values, subsequent receives will return zero values.
 ```
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	close(ch) // Closing a buffered channel

	// Trying to send on a closed channel will cause a panic
	// go func() {
	// 	ch <- 4
	// }()

	// Receiving from a closed buffered channel
	for {
		val, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(val) // Outputs: 1, 2, 3
	}

	// Further receives after exhausting values will yield zero values
	val, ok := <-ch
	fmt.Println(val, ok) // Outputs: 0 false
}

```





- To check if a channel is closed, use `close`: `val, ok := <-ch`. If closed, `ok` will be false.

- Alternatively, use range over a channel. When a channel closes, the range loop ends.

- Closing a closed channel, nil channel, or receive-only channel causes panic. Only a bidirectional channel or send-only channel can be closed.

- Closing a channel is not mandatory and not relevant for the garbage collector (GC). If a GC determines that a channel is unreachable, irrespective of whether it is open/closed, it will be garbage collected. Also sending value to unbuffer channel without go routine cause panic, since received not ready to receive.
- Below example causes panic. But if we replace with buffer of size 1 it will work.

```
        uch := make(chan int)

	uch <- 5

	close(uch)

	val, ok := <-uch
```


### Handling Concurrent Tasks in Go with Channels

In Go programming, when using goroutines (lightweight threads), they execute simultaneously, introducing the risk of race conditions when accessing shared resources. Consider a cashier scenario, where only a limited number of orders can be processed each day (e.g., 10 orders). Without coordination, concurrent access to the variable representing processed orders can lead to issues like slowing down, failure, or unexpected outcomes.

### Example: Cashier Processing Orders

In the example, multiple goroutines attempt to access the variable `ordersProcessed` with a limit of 10, potentially causing a race condition. Handling this scenario has various solutions:

1. **Increase the Limit:**
  - Increments the order processing limit, but with limits.

2. **Increase the Number of Cashiers:**
  - Introduces more resources for concurrent task handling.

3. **Distributed Work without Channels:**
  - However, this approach may still result in race conditions.

### The Role of Channels

To solve race conditions, channels provide a communication mechanism between goroutines, synchronizing their activities and eliminating the need for explicit locks (mutexes). This abstraction simplifies concurrent task coordination.

**Using Channels:**
1. **Create a Channel:**
  - Establish a channel to accept data for processing.

2. **Launch Goroutines:**
  - Start goroutines waiting for data on the channel.

3. **Pass Data Through the Channel:**
  - Send data to the channel through the main or other goroutines.

4. **Process Data Concurrently:**
  - Goroutines listening to the channel can accept and process data concurrently.

