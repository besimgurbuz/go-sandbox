package patterns

import (
	"fmt"
	"math/rand"
	"time"
)

// Generator: Function that returns a channel
func Boring(msg string) <-chan string { // Returns receive-only channel of strings
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

func ConcurrencyPatternMain1() {
	joe := Boring("Joe")
	ann := Boring("Ann")

	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}

	fmt.Println("You're both boring; I'm leaving.")
}

/*
* There is a important point, for example ann channel received something but if joe channel is not ann will wait to joe's reseive and execute

* We can solve this by multiplexing our channel with a wrapper channel
 */

func FanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

//* NOw we can use fanIn to listen them at the same time
func ConcurrencyPatternMain2() {
	c := FanIn(Boring("Joe"), Boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You're both boring; I'm leaving.")
}

// FanIn with select

func FanInWithSelect(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
}

// Timeout using select
func Timeout() {
	c := Boring("Joe")

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(100 * time.Millisecond):
			fmt.Println("You're too slow.")
			return
		}
	}
}

// Single timeout
func SingleTimeout() {
	c := Boring("Joe")
	timeout := time.After(200 * time.Millisecond)

	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}
}
