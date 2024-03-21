package main

import "fmt"

func qs(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}

	pivotIndex := partition(arr, lo, hi)

	qs(arr, lo, pivotIndex-1) // Handling the less/left side of the array excluding the pivot.
	qs(arr, pivotIndex+1, hi) // Handling the greater/right side of the array excluding the pivot
}

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi] // Random and can result in the worst case for time complexity
	index := lo - 1  // Used to swap the elements

	for i := lo; i < hi; i++ {
		if arr[i] < pivot {
			/*
				example:
					[8, 7, 6, 4, 5] where 5 is the pivot, index is 0 (FIRST LINE AFTER THE COMMENT)
					[4, 7, 6, 8, 5] 4 change position with 8
			*/
			index++
			tmp := arr[i]
			arr[i] = arr[index]
			arr[index] = tmp
		}
	}
	/*
		After throwing everything less than the pivot
		Increment the index value (to not change a value less than the pivot value)
		Change the pivot (remember that pivot := arr[hi]) to the element in the index
		And place the pivot in the "middle" (all values less in the left and all values greater in the right)
		USING THE EXAMPLE
		[4, 5, 6, 8, 7] 7 change position with pivot value (5)
	*/
	index++
	arr[hi] = arr[index]
	arr[index] = pivot
	return index
}

func quickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}

func main() {
	data := []int{9, 3, 7, 4, 69, 420, 42}
	quickSort(data)
	fmt.Println("Data should be [3, 4, 7, 9, 42, 69, 420]: ", data)
}
