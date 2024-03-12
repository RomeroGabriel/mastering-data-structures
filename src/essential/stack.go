package main

import "fmt"

type Node[T any] struct {
	value T
	prev  *Node[T]
}

type Stack[T any] struct {
	length int
	head   *Node[T]
}

func (s *Stack[T]) push(item T) {
	node := Node[T]{value: item}
	s.length++
	if s.head == nil {
		s.head = &node
		return
	}
	node.prev = s.head
	s.head = &node
}

func (s *Stack[T]) pop() *T {
	if s.length == 0 {
		return nil
	}
	s.length--

	oldHead := s.head
	s.head = oldHead.prev
	return &oldHead.value
}

func (s *Stack[T]) peek() *T {
	if s.head == nil {
		return nil
	}
	return &s.head.value
}

func main() {
	list := Stack[int]{}

	list.push(5)
	list.push(7)
	list.push(9)

	fmt.Println("Pop 9: ", *list.pop())
	fmt.Println("Legth: ", list.length)

	list.push(11)
	fmt.Println("Pop 11: ", *list.pop())
	fmt.Println("Pop 7: ", *list.pop())
	fmt.Println("Peek 5: ", *list.peek())
	fmt.Println("Pop 5: ", *list.pop())
	fmt.Println("Pop nothing: ", list.pop())

	list.push(70)
	fmt.Println("Peek 70: ", *list.peek())
	fmt.Println("Length: ", list.length)
}
