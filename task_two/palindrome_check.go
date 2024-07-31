package main

import (
	"fmt"
)

func palindromChecker(word string) {

	word = trimWord(word)
	word = toLowerCase(word)
	word = removePunctuation(word)

	if isPalindrome(word) {
		fmt.Println("The word ", word, " is a palindrome.")
	} else {
		fmt.Println("The word ", word, " is not palindrome.")
	}
}
