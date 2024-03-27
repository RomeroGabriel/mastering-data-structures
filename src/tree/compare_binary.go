package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func compare(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if a.value != b.value {
		return false
	}

	return compare(a.left, b.left) && compare(a.right, b.right)
}

func main() {
	a := &Node{value: 5}
	a.left = &Node{value: 3}
	a.right = &Node{value: 7}

	a.left.left = &Node{value: 1}
	a.left.right = &Node{value: 4}

	a.right.left = &Node{value: 6}
	a.right.right = &Node{value: 8}

	b := &Node{value: 5}
	b.left = &Node{value: 3}
	b.right = &Node{value: 7}

	b.left.left = &Node{value: 1}
	b.left.right = &Node{value: 4}

	b.right.left = &Node{value: 6}
	b.right.right = &Node{value: 8}

	fmt.Println("A and B is equal? ", compare(a, b))
	b.right.left.value = 9
	fmt.Println("A and B is equal? ", compare(a, b))
}
