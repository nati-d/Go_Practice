package main

import (
	"fmt"
)

func wordCounterFunction() {
	fmt.Print("Enter a sentence: ")
	word := acceptInput()


	

	word = trimWord(word)
	word = toLowerCase(word)
	word = removePunctuation(word)

	wordCount := wordCounter(word)

	fmt.Println("Word Count: ", wordCount)
}
