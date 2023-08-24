package main

import (
	"fmt"
	"sort"
	"strings"
)

func findAnagramSets(words *[]string) *map[string][]string {
	anagramSets := make(map[string][]string)

	for _, word := range *words {
		sortedWord := sortString(strings.ToLower(word))

		anagramSets[sortedWord] = append(anagramSets[sortedWord], word)
	}

	for sortedWord, words := range anagramSets {
		if len(words) == 1 {
			delete(anagramSets, sortedWord)
		}
	}

	return &anagramSets
}

func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"}
	anagramSets := findAnagramSets(&words)

	for _, words := range *anagramSets {
		fmt.Println(words)
	}
}
