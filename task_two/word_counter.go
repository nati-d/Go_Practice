package main

import (
	"fmt"
)

func wordCounterFunction(word string) {

	word = trimWord(word)
	word = toLowerCase(word)
	word = removePunctuation(word)

	wordCount := wordCounter(word)

	fmt.Println("Word Count: ", wordCount)
}
