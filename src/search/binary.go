package main

import (
	"fmt"
	"math"
)

func binarySearch(data []int, searchValue int) int {
	lo := 0
	hi := len(data)

	for lo < hi {
		middle := int(math.Floor(float64(lo + (hi-lo)/2)))
		value := data[int(middle)]

		if value == searchValue {
			return middle
		} else if value > searchValue {
			hi = middle
		} else {
			lo = middle + 1
		}
	}
	return -1
}

func main() {

	data := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	fmt.Println("Result for 69: ", binarySearch(data, 69))
	fmt.Println("Result for 1336: ", binarySearch(data, 1336))
	fmt.Println("Result for 69420: ", binarySearch(data, 69420))
	fmt.Println("Result for 69421: ", binarySearch(data, 69421))
	fmt.Println("Result for 1: ", binarySearch(data, 1))
	fmt.Println("Result for 0: ", binarySearch(data, 0))
}
