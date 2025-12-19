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

	node1.next = node2
	node2.next = node3
	node3.next = node2 //cycle
	fmt.Println(detectCycleEntry(node1))

}

func detectCycleEntry(node *ListNode) *ListNode {
	slow := node
	fast := node

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if fast == slow {
			entrypoint := node
			for entrypoint != slow {
				entrypoint = entrypoint.next
				slow = slow.next
			}
			return entrypoint
		}
	}
	return nil
}
