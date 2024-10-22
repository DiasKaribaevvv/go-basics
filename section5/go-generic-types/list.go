package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Insert(value T) {
	// Traverse to the last element of the list
	current := l
	for current.next != nil {
		current = current.next
	}

	// Insert new element at the end
	current.next = &List[T]{val: value}
}

func (l *List[T]) Print() {
	current := l
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}

func main() {
	var head *List[int] = &List[int]{val: 1}
	head.Insert(5)
	head.Print()

}
