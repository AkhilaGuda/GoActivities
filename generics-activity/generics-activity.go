package main

import "fmt"

type Queue[T any] struct { // T type parameter, queue can store any type
	items []T
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}
func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}
func (q *Queue[T]) Peek() (T, bool) {
	if len(q.items) == 0 {
		var zero T
		return zero, false
	}
	item := q.items[0]
	return item, true
}
func main() {
	var integerQueue Queue[int]
	integerQueue.Enqueue(1)
	integerQueue.Enqueue(2)
	integerQueue.Enqueue(3)
	integerQueue.Enqueue(4)
	fmt.Println("Elements in queue: ", integerQueue)
	value, flag := integerQueue.Dequeue()
	if !flag {
		fmt.Println("Queue is empty")
	} else {
		fmt.Println("Dequeue:", value)
	}
	value, flag = integerQueue.Peek()
	if !flag {
		fmt.Println("Queue is empty")
	} else {
		fmt.Println("Peek :", value)
	}

	var stringQueue Queue[string]
	stringQueue.Enqueue("trainees")
	stringQueue.Enqueue("akhila")
	stringQueue.Enqueue("guda")
	stringQueue.Enqueue("abcd")
	stringQueue.Enqueue("efgh")
	fmt.Println("Elements in queue: ", stringQueue)
	val, ok := stringQueue.Dequeue()
	if !ok {
		fmt.Println("Queue is empty")
	} else {
		fmt.Println("Dequeue:", val)
	}
	val, ok = stringQueue.Peek()
	if !ok {
		fmt.Println("Queue is empty")
	} else {
		fmt.Println("Peek :", val)
	}
}
