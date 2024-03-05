package main

import "fmt"

func bubbleSort(data *[7]int) {

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				tmp := data[j]
				data[j] = data[j+1]
				data[j+1] = tmp
			}
		}
	}
}

func main() {
	data := [7]int{9, 3, 7, 4, 69, 420, 42}
	fmt.Println("Inital data: ", data)
	bubbleSort(&data)
	fmt.Println("Final data: ", data)
}
