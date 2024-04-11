package main

import (
	"fmt"
	"slices"
)

type MinHeap struct {
	length int
	data   []int
}

/*
					50 (0)
		71 (1)						100 (2)
101(3)		80(4)			200(5)			101(6)
*/

func (h *MinHeap) Insert(value int) {
	h.data = slices.Grow(h.data, 1)
	h.data = append(h.data, value)
	h.heapifyUp(h.length)
	h.length++
}

func (h *MinHeap) Delete() int {
	if h.length == 0 {
		return -1
	}

	out := h.data[0]
	h.length--

	if h.length == 0 {
		h.data = []int{}
		return out
	}

	h.data[0] = h.data[h.length]
	h.heapifyDown(0)
	return out
}

func (h *MinHeap) heapifyDown(idx int) {
	if idx >= h.length {
		return
	}
	lIdx := h.getLeftChild(idx)
	rIdx := h.getRightChild(idx)

	if lIdx >= h.length || rIdx >= h.length {
		return
	}

	lV := h.data[lIdx]
	rV := h.data[rIdx]
	v := h.data[idx]

	// right value is the smallest, swap and heapifyDown
	if lV > rV && v > rV {
		h.data[idx] = rV
		h.data[rIdx] = v
		h.heapifyDown(rIdx)
	} else if rV > lV && v > lV {
		// left value is the smallest, swap and heapifyDown
		h.data[idx] = lV
		h.data[lIdx] = v
		h.heapifyDown(lIdx)
	}
}

func (h *MinHeap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}

	p := h.getParent(idx)
	parentV := h.data[p]
	idxValue := h.data[idx]

	// Go UP because is a MinHeap
	if parentV > idxValue {
		// Change value of idx with the parent
		h.data[idx] = parentV
		h.data[p] = idxValue
		h.heapifyUp(p)
	}
}

func (h *MinHeap) getParent(idx int) int {
	/*
		[50, 71, 100, 101, 80, 200, 101]
		[ 0,  1,   2,   3,  4,   5,   6]

			  EXAMPLE:
				PARENT OF 80:
				(4 - 1) / 2 = 1
	*/
	return (idx - 1) / 2
}

func (h *MinHeap) getLeftChild(idx int) int {
	/*
		[50, 71, 100, 101, 80, 200, 101]
		[ 0,  1,   2,   3,  4,   5,   6]

			EXAMPLE:
				LEFT OF 100:
				(2 * 2) + 1 = 5
	*/
	return idx*2 + 1
}

func (h *MinHeap) getRightChild(idx int) int {
	/*
		[50, 71, 100, 101, 80, 200, 101]
		[ 0,  1,   2,   3,  4,   5,   6]

			EXAMPLE:
				RIGHT OF 71:
				(1 * 2) + 2 = 4
	*/
	return idx*2 + 2
}

func main() {

	heap := MinHeap{}

	fmt.Println("Heap length: ", heap.length)

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(69)
	heap.Insert(420)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(7)

	fmt.Println("heap.length == 8: ", heap.length)
	fmt.Println("heap delete return 1 (smallest): ", heap.Delete())
	fmt.Println("heap delete return 3:", heap.Delete())
	fmt.Println("heap delete return 4:", heap.Delete())
	fmt.Println("heap delete return 5:", heap.Delete())
	fmt.Println("heap.length == 4: ", heap.length)
	fmt.Println("heap delete return 7:", heap.Delete())
	fmt.Println("heap delete return 8:", heap.Delete())
	fmt.Println("heap delete return 69:", heap.Delete())
	fmt.Println("heap delete return 420:", heap.Delete())
	fmt.Println("heap.length == 0: ", heap.length)

}
