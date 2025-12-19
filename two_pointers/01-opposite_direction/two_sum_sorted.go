package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(twosumsorted(nums, 6))
}

func twosumsorted(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			fmt.Println(nums[left], nums[right])
			return true
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return false
}
