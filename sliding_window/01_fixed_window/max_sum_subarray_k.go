package main

import "fmt"

func main() {
	nums := []int{2, 1, 5, 1, 3, 2}
	k := 3
	fmt.Println(maxSumarrayK(nums, k))
}

func maxSumarrayK(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}
	windowSum := 0
	maxSum := 0

	for i := 0; i < k; i++ {
		windowSum = windowSum + nums[i]
	}

	for i := k; i < len(nums); i++ {
		windowSum = windowSum + nums[i]   // add right element
		windowSum = windowSum - nums[i-k] // remove left element
		if windowSum > maxSum {
			maxSum = windowSum
		}
	}
	return maxSum
}
