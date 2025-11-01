package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countWords(text string) map[string]int {
	result := make(map[string]int)
	words := strings.Fields(text)
	for _, word := range words {
		word = strings.ToLower(word)
		word = strings.TrimFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r) && !unicode.IsNumber(r)
		})
		if word != "" {
			result[word]++
		}
	}
	return result
}

func main() {
	text := "Hello, world! Hello."
	counts := countWords(text)
	fmt.Println("Слово — Количество")
	for word, count := range counts {
		fmt.Printf("%s — %d\n", word, count)
	}
}
