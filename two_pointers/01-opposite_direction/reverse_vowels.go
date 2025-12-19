package main

import "fmt"

func main() {
	fmt.Println(reverseVowels("hello"))
	fmt.Println(reverseVowels("leetcode"))
}

func reverseVowels(str string) string {
	vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true}

	runes := []rune(str)

	left := 0
	right := len(runes) - 1

	for left < right {
		for left < right && !vowels[runes[left]] {
			left++
		}
		for left < right && !vowels[runes[right]] {
			right--
		}
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}
