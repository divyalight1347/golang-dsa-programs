package main

import "fmt"

func main() {
	fmt.Println(subarrayDivByK([]int{4, 5, 0, -2, -3, 1}, 5))
}

func subarrayDivByK(nums []int, k int) int {
	count := 0
	sum := 0

	remainderCount := map[int]int{0: 1}

	for _, num := range nums {
		sum = sum + num

		//handle negative numbers
		remainder := (sum%k + k) % k

		if freq, ok := remainderCount[remainder]; ok {
			count = count + freq
		}
		remainderCount[remainder]++
	}
	return count
}
