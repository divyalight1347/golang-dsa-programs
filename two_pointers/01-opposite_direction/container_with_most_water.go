package main

import "fmt"

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	result := maxArea(height)
	fmt.Println("Maximum water container area: ", result)
}

func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	maxWater := 0

	for left < right {
		width := right - left
		h := min(heights[left], heights[right])
		area := h * width
		maxWater = max(maxWater, area)
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return maxWater
}
