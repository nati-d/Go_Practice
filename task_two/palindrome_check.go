package main

import (
	"bufio"
	"fmt"
	"os"
)

func palindromChecker() {
	fmt.Print("Enter a word: ")
	var reader = bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString('\n')

	

	word = trimWord(word)
	word = toLowerCase(word)
	word = removePunctuation(word)

	if isPalindrome(word) {
		fmt.Println("The word ", word, " is a palindrome.")
	} else {
		fmt.Println("The word ", word, " is not palindrome.")
	}
}
