package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func PreOrderSearch(curr *Node, path *[]int) {
	if curr == nil {
		return
	}
	*path = append(*path, curr.value)
	PreOrderSearch(curr.left, path)
	PreOrderSearch(curr.right, path)
}

func InOrderSearch(curr *Node, path *[]int) {
	if curr == nil {
		return
	}
	PreOrderSearch(curr.left, path)
	*path = append(*path, curr.value)
	PreOrderSearch(curr.right, path)
}

func PostOrderSearch(curr *Node, path *[]int) {
	if curr == nil {
		return
	}
	PreOrderSearch(curr.left, path)
	PreOrderSearch(curr.right, path)
	*path = append(*path, curr.value)
}

func main() {
	head := Node{value: 5}
	head.left = &Node{value: 3}
	head.right = &Node{value: 7}

	head.left.left = &Node{value: 1}
	head.left.right = &Node{value: 4}

	head.right.left = &Node{value: 6}
	head.right.right = &Node{value: 8}

	result := make([]int, 0, 10)
	PreOrderSearch(&head, &result)
	fmt.Println("PreOrder: ", result)
	result = make([]int, 0, 10)

	InOrderSearch(&head, &result)
	fmt.Println("InOrder: ", result)
	result = make([]int, 0, 10)

	PostOrderSearch(&head, &result)
	fmt.Println("PostOrder: ", result)
}
