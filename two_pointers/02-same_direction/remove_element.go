package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	k := removeElement(nums, 2)
	fmt.Println(nums[:k])
}

func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
