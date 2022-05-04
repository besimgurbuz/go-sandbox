package patterns

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web1   = fakeSearch("web")
	Web2   = fakeSearch("web")
	Image1 = fakeSearch("image")
	Image2 = fakeSearch("image")
	Video1 = fakeSearch("video")
	Video2 = fakeSearch("video")
)

type Result string
type Search func(query string) Result

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}

func Google(query string) (results []Result) {
	c := make(chan Result)

	go func() {
		c <- Web1(query)
	}()
	go func() {
		c <- Image1(query)
	}()
	go func() {
		c <- Video1(query)
	}()

	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func FakeGoogleMain() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)

	fmt.Println(results)
	fmt.Println(elapsed)
}

// How to avoid timeout
// * Replicate the servers. Send requests to multiple replicas, and use the first response.
func First(query string, replicas ...Search) Result {
	c := make(chan Result)
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}

	return <-c
}

func GoogleWithReplica(query string) (results []Result) {
	c := make(chan Result)

	go func() {
		c <- First(query, Web1, Web2)
	}()
	go func() {
		c <- First(query, Image1, Image2)
	}()
	go func() {
		c <- First(query, Video1, Video2)
	}()
	timeout := time.After(80 * time.Millisecond)

	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func FakeGoogleMainReplicatedExample() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := First(
		"golang",
		fakeSearch("replica 1"),
		fakeSearch("replica 2"))
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)
}

func FakeGoogleReplicaMain() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := GoogleWithReplica("golang")
	elapsed := time.Since(start)

	fmt.Println(results)
	fmt.Println(elapsed)
}
