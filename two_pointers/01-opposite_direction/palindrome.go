package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("madam"))
	fmt.Println(isPalindrome("Divya"))
}

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
