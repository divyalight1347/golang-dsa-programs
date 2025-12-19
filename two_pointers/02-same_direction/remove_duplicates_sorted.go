package main

import "fmt"

func main() {
	nums := []int{2, 2, 5, 5, 6, 7, 8, 8, 9, 9}
	k := removeDuplicatesSorted(nums)
	fmt.Println(nums[:k])
}

func removeDuplicatesSorted(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
