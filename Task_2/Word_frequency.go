package main

import (
	"fmt"
	"strings"
	"unicode"
)

func wordFrequencyCount(input string) map[string]int {
	wordCount := make(map[string]int)
	words := strings.FieldsFunc(input, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})
	
	for _, word := range words {
		word = strings.ToLower(word)
		wordCount[word]++
	}
	
	return wordCount
}

func main() {
	input := "Hello, hello! How are you? You are learning Go, Go is great."
	frequency := wordFrequencyCount(input)
	fmt.Println(frequency)
}
