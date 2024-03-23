package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Node[T any] struct {
	value T
	prev  *Node[T]
	next  *Node[T]
}

type DoublyLinkedList[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func (dll *DoublyLinkedList[T]) Prepend(data T) {
	node := Node[T]{value: data}

	dll.length++
	if dll.head == nil {
		dll.head, dll.tail = &node, &node
		return
	}

	/* Example: Prepend Z
	A <-> B <-> C, HEAD = A
	1. Z.NEXT = A
	2. A.PREV = Z
	3. DLL.HEAD = Z
	Z <-> A <-> B <-> C, HEAD = Z
	*/
	node.next = dll.head  // 1
	dll.head.prev = &node // 2
	dll.head = &node      // 3
}

func (dll *DoublyLinkedList[T]) InsertAt(data T, index int) error {
	if index > dll.length {
		return errors.New("index bigger than list size")
	}
	if index == dll.length {
		// Is just add in the end
		dll.Append(data)
		return nil
	}
	if index == 0 {
		// Is just add in the beginning
		dll.Prepend(data)
		return nil
	}
	/* Example: insert Z at 2
	A <-> B <-> C, HEAD = A
	1. Z.NEXT = C
	2. Z.PREV =  B
	3. C.PREV = Z
	4. B.NEXT = Z
	A <-> B <-> Z <-> C
	*/
	curr := dll.head
	for i := 0; curr != nil && i < index; i++ {
		curr = curr.next
	}
	dll.length++
	node := Node[T]{value: data}
	node.next = curr       // 1
	node.prev = curr.prev  // 2
	curr.prev = &node      // 3
	node.prev.next = &node // 4
	return nil
}

func (dll *DoublyLinkedList[T]) Append(data T) {
	node := Node[T]{value: data}
	dll.length++
	if dll.tail == nil {
		dll.head, dll.tail = &node, &node
		return
	}
	/* Example: insert Z
	A <-> B <-> C, HEAD = A
	1. Z.PREV = C
	2. C.NEXT = Z
	3. DLL.TAIL = Z
	A <-> B <-> C <-> Z
	*/
	node.prev = dll.tail
	dll.tail.next = &node
	dll.tail = &node
}

func (dll *DoublyLinkedList[T]) Remove(data T) {
	curr := dll.head
	dataComparable := reflect.ValueOf(data)
	if !dataComparable.Comparable() {
		panic("type cannot be compared")
	}
	for i := 0; curr != nil && i < dll.length; i++ {
		currReflectValue := reflect.ValueOf(curr.value)
		if currReflectValue.Equal(dataComparable) {
			break
		}
		curr = curr.next
	}
	if curr == nil {
		return
	}
	dll.removeNode(curr)
}

func (dll *DoublyLinkedList[T]) RemoveAt(index int) {
	node := dll.GetAt(index)
	if node == nil {
		return
	}
	dll.removeNode(node)
}

func (dll *DoublyLinkedList[T]) removeNode(node *Node[T]) {
	dll.length--
	if dll.length == 0 {
		dll.head, dll.tail = nil, nil
		return
	}
	if node == dll.head {
		dll.head = node.next
	}

	if node == dll.tail {
		dll.tail = node.prev
	}
	/* Example: delete B
	A <-> B <-> C, HEAD = A
	1. A.NEXT = C
	2. C.PREV = A
	A <-> C
	*/
	if node.prev != nil {
		node.prev.next = node.next // 1
	}
	if node.next != nil {
		node.next.prev = node.prev // 2
	}

}

func (dll *DoublyLinkedList[T]) Get(index int) *T {
	node := dll.GetAt(index)
	if node == nil {
		return nil
	}
	return &node.value
}

func (dll *DoublyLinkedList[T]) GetAt(index int) *Node[T] {
	curr := dll.head
	for i := 0; curr != nil && i < index; i++ {
		curr = curr.next
	}
	return curr
}

func main() {
	list := &DoublyLinkedList[int]{}

	list.Append(1)
	list.Prepend(2)
	fmt.Println("Head value: ", list.head.value)
	fmt.Println("Tail value: ", list.tail.value)

	list.InsertAt(5, 1)
	list.Append(6)
	fmt.Println("Get value from 2: ", list.GetAt(2).value)

	fmt.Println("Printing all list: ")
	head := list.head
	for head != nil {
		fmt.Println("	", head.value)
		head = head.next
	}

	list.Remove(5)
	list.RemoveAt(0)
	fmt.Println("Head value after delete index 0: ", list.head.value)
	list.RemoveAt(list.length)
	fmt.Println("Tail value after delete list lenght: ", list.tail.value)

	fmt.Println("Printing all list: ")
	head = list.head
	for head != nil {
		fmt.Println("	", head.value)
		head = head.next
	}
}
