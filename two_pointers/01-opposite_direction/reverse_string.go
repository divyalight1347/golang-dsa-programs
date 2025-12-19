package main

import "fmt"

func main() {
	str := "Divya"
	runes := []rune(str)
	reverseString(runes)
	fmt.Println(string(runes))
}

func reverseString(str []rune) {
	left := 0
	right := len(str) - 1

	for left < right {
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
}
