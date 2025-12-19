package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {
	slow := n
	fast := getNext(n)

	for fast != 1 && slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}
	return fast == 1
}

func getNext(n int) int {
	sum := 0

	for n > 0 {
		digit := n % 10
		sum = sum + digit*digit
		n = n / 10
	}
	return sum
}
