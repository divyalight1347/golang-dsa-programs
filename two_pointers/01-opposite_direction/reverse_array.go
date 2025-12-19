package main

import "fmt"

func main() {
	nums := []int{6, 7, 8, 4, 7}
	reverseArray(nums)
	fmt.Println(nums)
}

func reverseArray(nums []int) {
	left := 0
	right := len(nums) - 1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--

	}
}
