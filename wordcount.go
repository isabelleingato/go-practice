package main

import (
	"strings"
)

func WordCount(s string) map[string]int {
	var elem int
	words := strings.Fields(s)
	wordsMap := make(map[string]int, len(words))
	for _, value := range words {
		elem = wordsMap[value]
		wordsMap[value] = elem + 1
	}
	return wordsMap
}

