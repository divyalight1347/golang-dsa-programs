package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

func main() {
	n1 := &ListNode{1, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{3, nil}
	n4 := &ListNode{4, nil}
	n5 := &ListNode{5, nil}

	n1.next = n2
	n2.next = n3
	n3.next = n4
	n4.next = n5
	fmt.Println(middleInList(n1))

}

func middleInList(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow
}
