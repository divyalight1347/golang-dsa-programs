package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func main() {
	n1 := &ListNode{val: 1}
	n2 := &ListNode{val: 2}
	n3 := &ListNode{val: 2}
	n4 := &ListNode{val: 1}
	n1.next = n2
	n2.next = n3
	n3.next = n4

	fmt.Println(isPalindrome(n1))
}

func isPalindrome(node *ListNode) bool {
	if node == nil || node.next == nil {
		return true
	}

	slow := node
	fast := node

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	var prev *ListNode

	for slow != nil {
		next := slow.next
		slow.next = prev
		prev = slow
		slow = next
	}
	left := node
	right := prev

	for right != nil {
		if left.val != right.val {
			return false
		}
		left = left.next
		right = right.next
	}
	return true

}
