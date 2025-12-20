package main

import "fmt"

func main() {
	fmt.Println(numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3))
}

func numberOfSubarrays(nums []int, k int) int {
	count := 0
	sum := 0

	prefixSumCount := map[int]int{0: 1}

	for _, num := range nums {
		if num%2 != 0 {
			sum = sum + 1
		}
		if freq, ok := prefixSumCount[sum-k]; ok {
			count = count + freq
		}
		prefixSumCount[sum]++
	}
	return count
}
