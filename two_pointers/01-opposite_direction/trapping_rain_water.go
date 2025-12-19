package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trapingWater(nums))
}

func trapingWater(nums []int) int {
	left := 0
	right := len(nums) - 1
	leftMax := 0
	rightMax := 0
	water := 0

	for left < right {
		if nums[left] < nums[right] {
			if nums[left] >= leftMax {
				leftMax = nums[left]
			} else {
				water = water + leftMax - nums[left]
			}
			left++
		} else {
			if nums[right] >= rightMax {
				rightMax = nums[right]
			} else {
				water = water + rightMax - nums[right]
			}
			right--
		}
	}
	return water
}
