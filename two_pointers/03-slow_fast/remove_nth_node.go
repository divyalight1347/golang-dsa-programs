package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func main() {
	n1 := &ListNode{val: 1}
	n2 := &ListNode{val: 2}
	n3 := &ListNode{val: 3}

	n1.next = n2
	n2.next = n3
	head := removeNthNode(n1, 1)
	for head != nil {
		fmt.Println(head.val, " ")
		head = head.next
	}

}

func removeNthNode(head *ListNode, n int) *ListNode {
	dummy := &ListNode{next: head}
	slow := dummy
	fast := dummy

	for i := 0; i < n; i++ {
		fast = fast.next
	}

	for fast.next != nil {
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
	return dummy.next
}
