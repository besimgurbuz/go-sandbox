package concurrency

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func walkRecurse(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkRecurse(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walkRecurse(t.Right, ch)
	}
}

func Walk(t *tree.Tree, ch chan int) {
	walkRecurse(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	firstTreeCh := make(chan int)
	secondTreeCh := make(chan int)

	go Walk(t1, firstTreeCh)
	go Walk(t2, secondTreeCh)

	for {
		val1, ok1 := <-firstTreeCh
		val2, ok2 := <-secondTreeCh

		if val1 != val2 {
			return false
		}

		if !ok1 && !ok2 {
			break
		}
	}

	return true
}

func ConcurrencyExercise() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

	fmt.Println("Should be true", Same(tree.New(1), tree.New(1)))
	fmt.Println("Should be false", Same(tree.New(1), tree.New(2)))
}
