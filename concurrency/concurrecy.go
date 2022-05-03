package concurrency

import (
	"fmt"
	"time"
)

func Say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}

/*
	Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`

		ch <- v // send v to channel ch
		v := <-ch // Receive from ch, and assign value to v.
	(the data flows in the direction of the arrow)

	Like maps and slices, channels must be created before use:
		ch := make(chan int)

	By default, sends and receives block until the other side is ready. This allows goroutines to synchronize wihtout
	explicit locks or condition values.
*/

func Sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send tum to c
}

func SumMain() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go Sum(s[:len(s)/2], c)
	go Sum(s[len(s)/2:], c)

	x, y := <-c, <-c // receive from c
	fmt.Println("This should wait")
	fmt.Println(x, y, x+y)
}

/*
	Buffered Channels
	Channels can be buffered. Provide the buffer length as the second argument to make to initialzie a buffered channel:
		ch := make(chan int, 100)

	Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
*/

func BufferedMain() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
	A sender can close a chennel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

		v, is_open := <-ch
		is_open is false if there are no more values to receive and channel is closed.

	the loop for i := range c receives values from the channel until it is closed

	Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

	Another Note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a arange loop.
*/

func Fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func FibonacciMain() {
	c := make(chan int, 10)
	go Fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

/*
	Select
	The select statement lets a goroutine wait on multiple communication operations.

	A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
*/

func FibonacciWithSelect(c, quit chan int) {
	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func FibonacciMainWithSelect() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	FibonacciWithSelect(c, quit)
}

/*
	Default Selection
	The default case in a select is run if no other case is ready.

	select {
	case i := <-c:
		 use i
	default:
		receiving from c would block
	}
*/

func DefaultSelectMain() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("      .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
