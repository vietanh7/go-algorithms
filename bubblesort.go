package main

import (
	"log"
)

var Log = log.Println

func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func main() {
	slice := []int{9, 3, 5, 4, 6, 8, 2, 7, 1}
	slice = bubbleSort(slice)
	Log("Sorted slice: ", slice)
}
