package main

import "fmt"

type NumArray struct {
	prefix []int
}

func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)

	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	return NumArray{prefix}
}

func (n *NumArray) SumRange(left, right int) int {
	return n.prefix[right+1] - n.prefix[left]
}
func main() {
	nums := []int{1, 2, 3, 4}
	obj := Constructor(nums)
	fmt.Println(obj.SumRange(1, 3))
}
