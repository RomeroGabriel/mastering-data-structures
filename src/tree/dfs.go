package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

func search(curr *TreeNode, searchValue int) bool {
	if curr == nil {
		return false
	}

	if curr.value == searchValue {
		return true
	}

	if curr.value < searchValue {
		return search(curr.right, searchValue)
	}
	return search(curr.left, searchValue)
}

func dfs(head *TreeNode, searchValue int) bool {
	return search(head, searchValue)
}

func main() {
	head := &TreeNode{value: 5}
	head.left = &TreeNode{value: 3}
	head.right = &TreeNode{value: 7}

	head.left.left = &TreeNode{value: 1}
	head.left.right = &TreeNode{value: 4}

	head.right.left = &TreeNode{value: 6}
	head.right.right = &TreeNode{value: 8}
	fmt.Println("Is 2 present in treeNode? ", dfs(head, 2))
	fmt.Println("Is 1 present in treeNode? ", dfs(head, 1))
	fmt.Println("Is 3 present in treeNode? ", dfs(head, 3))
	fmt.Println("Is 4 present in treeNode? ", dfs(head, 4))
	fmt.Println("Is 8 present in treeNode? ", dfs(head, 8))
	fmt.Println("Is 9 present in treeNode? ", dfs(head, 9))
}
