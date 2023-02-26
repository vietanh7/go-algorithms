package main

import (
	"log"
	"math"
)

var Log = log.Println

// Find max subarray
func maxSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	mid := len(nums) / 2
	leftMax := maxSubarray(nums[:mid])
	rightMax := maxSubarray(nums[mid:])
	crossMax := maxCrossingSubarray(nums, mid)

	return max(leftMax, rightMax, crossMax)
}

// Find max of crossing subarrays
func maxCrossingSubarray(nums []int, mid int) int {
	leftMax := math.MinInt64
	sum := 0
	for i := mid - 1; i >= 0; i-- {
		sum += nums[i]
		leftMax = max(leftMax, sum)
	}

	rightMax := math.MinInt64
	sum = 0
	for i := mid; i < len(nums); i++ {
		sum += nums[i]
		rightMax = max(rightMax, sum)
	}

	return leftMax + rightMax
}

// Find the maximum value among a variadic numbers
func max(nums ...int) int {
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	Log(maxSubarray(nums)) // Output: 6
}
