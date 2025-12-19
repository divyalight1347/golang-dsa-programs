package main

import "fmt"

func main() {
	fmt.Println(validPalindrome("aba"))
	fmt.Println(validPalindrome("abca"))
	fmt.Println(validPalindrome("abc"))
}

func validPalindrome(str string) bool {
	left := 0
	right := len(str) - 1

	for left < right {
		if str[left] != str[right] {
			return isPalindrome(str, left, right-1) || isPalindrome(str, left+1, right)
		}
		left++
		right--
	}
	return true
}

func isPalindrome(str string, left, right int) bool {
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
