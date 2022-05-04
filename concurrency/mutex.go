/*
	sync.Mutex

	We've seen how channels are great for communication among goroutines.

	But what if we don't need communication? What if we just want to make sure only one goroutine can access a variable at a a time to avoid conflicts?

	This concept is called mutual exclusion, and the convertional name for the data structure that provides it is mutex.

	Go's standard library provides mutual exclusion whith sync.Mutex and its two methods:

		Lock
		Unlock

	We can defina a block of code to be executed in mutual exclusion by surronding it with a call to Lock and Unlock as shown on the Inc method.

	We can also use defer to unsure the mutex will be unlocked as in the Value method.
*/
package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func MutexMain() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
