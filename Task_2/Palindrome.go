package main

import (
	"fmt"

	"github.com/mohaali/a2sv-backend-learning-path/task-2/pkg"
)

func main() {
	fmt.Println("                              Palindrome Check ✅/🚫")
	fmt.Println("⚠⚠This checker is case-insensitive, ignores spaces and ignores punctuation marks.⚠⚠")
	fmt.Println()
	fmt.Println()
	fmt.Printf("Enter your word: ")
	str := pkg.StringInput()
	fmt.Println()

	palindrome := pkg.PalindromeCheck(str)

	fmt.Println()
	fmt.Println()
	fmt.Printf("Your word is ")
	if palindrome {
		fmt.Printf("a palindrome ✅.")
	} else {
		fmt.Printf("not a palindrome 🚫.")

	}
	fmt.Println()
	fmt.Println()

}
