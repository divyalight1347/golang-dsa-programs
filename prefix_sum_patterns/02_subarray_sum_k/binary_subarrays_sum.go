package main

import "fmt"

func main() {
	fmt.Println(numSubarrayWithSum([]int{1, 0, 1, 0, 1}, 2))
}

func numSubarrayWithSum(nums []int, goal int) int {
	count := 0
	sum := 0

	prefixSumCount := map[int]int{0: 1}

	for _, num := range nums {
		sum = sum + num
		if freq, ok := prefixSumCount[sum-goal]; ok {
			count = count + freq
		}
		prefixSumCount[sum]++
	}
	return count
}
