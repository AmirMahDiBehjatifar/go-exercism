package main

import (
	"fmt"
	"strings"
)

func main() {
	// Examples That you can test // Output false
	// harward
	//IsIsoGram
	// Aftaby

	// Things that should return true
	// Emily Jung Schwartzkopf
	// lumberjacks
	// background
	// downstream
	// six-year-old
	fmt.Println(IsIsogram("Emily Jung Schwartzkopf"))

}
func IsIsogram(word string) bool {
	// if word be empty it would return true
	if word == "" {
		return true
	}
	wordFilter := strings.NewReplacer(" ", "", "-", "")
	word = wordFilter.Replace(word)
	word = strings.ToLower(word)

	var wordIndex int = len(word) - 1 // if len == 4 it return 3 because we count index from zero
	for i := 0; i < len(word)-1; i++ {

		// i developed an algorithm that works like this:
		// for example if the word length is 4 (4 - 1 = 3 cause we count index from 0)
		// First, check if word[3] matches any of loopBody that is contain [2,1,0].If not, decrement wordIndex
		// Then,check if word[2] matches any of LoopBody loop[1,0], and so on
		for x := 0; x < wordIndex; x++ {
			if word[wordIndex] == word[x] {
				return false
			}
		}
		wordIndex--
	}
	return true
}
