package main

import "fmt"

func main() {

	words := []string{"hi", "hello", "day", "hi", "night", "day"}

	frequencies := map[string]int{}

	frequencies = wordFrequencies(words)

	fmt.Println(frequencies)
}

func wordFrequencies(words []string) map[string]int {
	frequencies := map[string]int{}

	for _, word := range words {
		frequencies[word]++
	}

	return frequencies
}
