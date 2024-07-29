package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	fmt.Print("Enter a word: ")
	word, _ := reader.ReadString('\n')

	word = trimWord(word)
	word = toLowerCase(word)
	word = removePunctuation(word)

	if isPalindrome(word) {
		fmt.Println("The word " ,word, " is a palindrome.")
	} else {
		fmt.Println("The word " ,word, " is not palindrome.")
	}
}

func removePunctuation(word string) string {
	var result string
	for i := 0; i < len(word); i++ {
		if int(word[i]) >= 65 && int(word[i]) <= 90 || int(word[i]) >= 97 && int(word[i]) <= 122 {
			result += string(word[i])
		}
	}
	return result
}


func toLowerCase(word string) string {
	return strings.ToLower(word)
}



func trimWord(word string) string {
	return strings.Trim(word, " ")
}


func isPalindrome(word string) bool {
	var reversedWord string

	for i := len(word) - 1; i >= 0; i-- {
		reversedWord += string(word[i])
	}

	return word == reversedWord
}