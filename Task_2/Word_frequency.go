package main

import (
	"fmt"

	"github.com/mohaali/a2sv-backend-learning-path/task-2/pkg"
)

func main() {
	fmt.Println("                              Word Counter ⏱")
	fmt.Println("⚠⚠This counter is case-insensitive and ignores punctuation marks.⚠⚠")
	fmt.Println()
	fmt.Println()
	fmt.Printf("Enter your sentence: ")
	str := pkg.StringInput()
	fmt.Println()
	counter := pkg.CountWords(str)

	fmt.Println()
	fmt.Println("Word     Count")
	fmt.Println()
	for word, count := range counter {
		fmt.Printf("%-5v .......... %v\n", word, count)
	}
	fmt.Println()
	fmt.Println()

}
