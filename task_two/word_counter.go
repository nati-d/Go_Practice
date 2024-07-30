package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordCounterFunction() {
	fmt.Print("Enter a sentence: ")
	var reader = bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString('\n')


	

	word = trimWord(word)
	word = toLowerCase(word)
	word = removePunctuation(word)

	wordCount := wordCounter(word)

	fmt.Println("Word Count: ", wordCount)
}
