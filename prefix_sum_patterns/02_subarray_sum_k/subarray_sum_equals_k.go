package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 2, 3}, 3))
	fmt.Println(subarraySum([]int{1, -1, 0}, 0))
	fmt.Println(subarraySum([]int{3, 4, 7, 2 - 3}, 7))
}

func subarraySum(nums []int, k int) int {
	count := 0
	sum := 0

	prefixSumCount := map[int]int{0: 1}

	for _, num := range nums {
		sum = sum + num
		if freq, ok := prefixSumCount[sum-k]; ok {
			count = count + freq
		}
		prefixSumCount[sum]++
	}
	return count
}
