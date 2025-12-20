package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func main() {
	node1 := &ListNode{val: 1}
	node2 := &ListNode{val: 2}
	node3 := &ListNode{val: 3}
	node4 := &ListNode{val: 4}

	node1.next = node2
	node2.next = node3
	node3.next = node4

	reversed := reverseList(node1)

	for reversed != nil {
		fmt.Print(reversed.val, " ")
		reversed = reversed.next
	}

}

func reverseList(node *ListNode) *ListNode {
	var prev *ListNode
	curr := node
	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	return prev
}
