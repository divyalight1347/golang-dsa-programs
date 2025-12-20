package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	res := reverseSublist(head, 2, 4)

	for res != nil {
		fmt.Print(res.val, " ")
		res = res.next
	}
}

func reverseSublist(head *ListNode, left, right int) *ListNode {
	if head == nil || left == right {
		return head
	}

	dummy := &ListNode{next: head}
	prev := dummy
	// 1️⃣ Move prev to node before `left`
	for i := 1; i < left; i++ {
		prev = prev.next
	}
	// 2️⃣ Reverse sublist
	curr := prev.next

	var prevNode *ListNode

	for i := 0; i <= right-left; i++ {
		next := curr.next
		curr.next = prevNode
		prevNode = curr
		curr = next
	}

	// 3️⃣ Reconnect
	prev.next.next = curr //5 in this case
	prev.next = prevNode

	return dummy.next
}
