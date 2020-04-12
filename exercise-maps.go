package main

import (
	"golang.org/x/tour/wc"

	"strings"
)

func WordCount(s string) map[string]int {
	// Get input string as a slice of strings (words)
	var words []string = strings.Fields(s) 

	// Map to store words and their counts
	counts := make(map[string]int)

	for _, word := range words {
		_, ok := counts[word] // Check if word already in map
		if ok {
			counts[word] = counts[word] + 1

		} else {
			counts[word] = 1
		}
	}

	return counts
}

func main() {
	wc.Test(WordCount)
}
