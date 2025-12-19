package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 4, 0, 5, 2, 0}
	movezeroes(nums)
	fmt.Println(nums)
}

func movezeroes(nums []int) {
	slow := 0

	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}
