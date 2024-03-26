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
	if q.length == 0 {
		q.tail = nil
	}
	return &head.value
}

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func bfs(head TreeNode, searchValue int) bool {
	queue := Queue[TreeNode]{}
	queue.enqueue(head)
	for queue.length != 0 {
		curr := queue.deque()
		if curr == nil {
			continue
		}

		if curr.value == searchValue {
			return true
		}

		if curr.left != nil {
			queue.enqueue(*curr.left)
		}
		if curr.right != nil {
			queue.enqueue(*curr.right)
		}
	}
	return false
}

func main() {
	head := TreeNode{value: 5}
	head.left = &TreeNode{value: 3}
	head.right = &TreeNode{value: 7}

	head.left.left = &TreeNode{value: 1}
	head.left.right = &TreeNode{value: 4}

	head.right.left = &TreeNode{value: 6}
	head.right.right = &TreeNode{value: 8}
	fmt.Println("Is 2 present in treeNode? ", bfs(head, 2))
	fmt.Println("Is 1 present in treeNode? ", bfs(head, 1))
	fmt.Println("Is 3 present in treeNode? ", bfs(head, 3))
	fmt.Println("Is 4 present in treeNode? ", bfs(head, 4))
	fmt.Println("Is 8 present in treeNode? ", bfs(head, 8))
	fmt.Println("Is 9 present in treeNode? ", bfs(head, 9))
}
