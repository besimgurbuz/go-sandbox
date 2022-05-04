package list

import (
	"fmt"

	"errors"
)

// A type to create a linked list
type List[T any] struct {
	Next *List[T]
	Val  T
}

func (l *List[T]) Print() {
	index := 0
	for c := l; c != nil; c = c.Next {
		v := c.Val
		fmt.Printf("index: %d, value: %v\n", index, v)
		index++
	}
}

func (l *List[T]) Remove(item *List[T]) (*List[T], error) {
	var prevItem *List[T] = l

	if item == l {
		*l = *l.Next

		return l, nil
	}

	for c := l; c != nil; c = c.Next {
		if c == item {
			prevItem.Next = c.Next
			return c, nil
		} else {
			prevItem = c
		}
	}

	return nil, errors.New("item not found")
}

func (l *List[T]) RemoveByIndex(i int) (*List[T], error) {
	index := 0
	prevItem := l

	if index == i {
		*l = *l.Next

		return l, nil
	}

	for c := l; c != nil; c = c.Next {
		if index == i {
			prevItem.Next = c.Next

			return c, nil
		} else {
			index++
			prevItem = c
		}
	}

	return nil, errors.New("item not found")
}

func (l *List[T]) Add(newValue List[T], index int) {
	i := 0
	for curr := l; curr != nil; curr = curr.Next {
		if i == index {
			temp := curr.Next
			newValue.Next = temp
			curr.Next = &newValue
			break
		} else {
			i++
		}
	}
}

func (l *List[T]) ConvertToSlice() []T {
	var slice []T
	i := 0

	for curr := l; curr != nil; curr = curr.Next {
		slice = append(slice, curr.Val)
		i++
	}

	return slice
}
