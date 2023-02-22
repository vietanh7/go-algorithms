package main

import (
	"log"
)

var Log = log.Println

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	leftSortedArr := mergeSort(arr[:mid])
	rightSortedArr := mergeSort(arr[mid:])
	arr = merge(leftSortedArr, rightSortedArr)
	return arr
}

func merge(leftArr []int, rightArr []int) []int {
	mergedArr := make([]int, 0, len(leftArr)+len(rightArr))
	i, j := 0, 0
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] <= rightArr[j] {
			mergedArr = append(mergedArr, leftArr[i])
			i++
		} else {
			mergedArr = append(mergedArr, rightArr[j])
			j++
		}
	}
	mergedArr = append(mergedArr, leftArr[i:]...)
	mergedArr = append(mergedArr, rightArr[j:]...)
	return mergedArr
}

func main() {
	slice := []int{5, 3, 4, 6, 7, 9, 8, 2, 1}
	slice = mergeSort(slice)
	Log("Sorted slice: ", slice)
}
