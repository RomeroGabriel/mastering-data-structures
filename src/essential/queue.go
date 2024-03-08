package main

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Queue[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func (q *Queue[T]) enqueue(item T) {
	node := Node[T]{value: item}
	q.length++
	if q.tail == nil {
		// First element
		q.tail, q.head = &node, &node
		return
	}
	// Last item next now appoint to the new node
	q.tail.next = &node
	// Node became the last item
	q.tail = &node
}

func (q *Queue[T]) deque() *T {
	if q.head == nil {
		return nil
	}
	q.length--
	head := q.head
	//  change head to next element
	q.head = q.head.next
	head.next = nil
	return &head.value
}

func (q *Queue[T]) peek() *T {
	if q.head == nil {
		return nil
	}
	return &q.head.value
}

func (q *Queue[T]) rear() *T {
	if q.tail == nil {
		return nil
	}
	return &q.tail.value
}

func main() {
	testQueue := Queue[int]{}
	testQueue.enqueue(5)
	testQueue.enqueue(7)
	testQueue.enqueue(9)
	fmt.Println("Deque 5 (FIFO): ", *testQueue.deque())
	fmt.Println("Lenght: ", testQueue.length)
	testQueue.enqueue(11)
	fmt.Println("Deque 7 (FIFO): ", *testQueue.deque())
	fmt.Println("Deque 9 (FIFO): ", *testQueue.deque())
	fmt.Println("Peek 11: ", *testQueue.peek())
	fmt.Println("Deque 11 (FIFO): ", *testQueue.deque())
	fmt.Println("Deque nill: ", testQueue.deque())
	fmt.Println("Lenght: ", testQueue.length)
}
