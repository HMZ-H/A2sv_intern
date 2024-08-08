package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindrome(input string) bool {
	var sanitized []rune
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			sanitized = append(sanitized, unicode.ToLower(r))
		}
	}

	n := len(sanitized)
	for i := 0; i < n/2; i++ {
		if sanitized[i] != sanitized[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	input := "A man, a plan, a canal, Panama"
	fmt.Println(isPalindrome(input)) // Output: true
}
