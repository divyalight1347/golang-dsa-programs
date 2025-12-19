package main

import "fmt"

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSqaures(nums))
}

func sortedSqaures(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	left := 0
	right := n - 1
	pos := n - 1

	for left <= right {
		leftSq := nums[left] * nums[left]
		rightSq := nums[right] * nums[right]
		if leftSq > rightSq {
			result[pos] = leftSq
			left++
		} else {
			result[pos] = rightSq
			right--
		}
		pos--
	}
	return result
}
