package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("What Do You Want To Do?")
	fmt.Println("1. Count Words")
	fmt.Println("2. Check Palindrome")
	var choice int
	fmt.Print("Enter Your Choice: ")
	fmt.Scanln(&choice)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a word: ")
	sentence, _ := reader.ReadString('\n')
	sentence = trimWord(sentence)
	sentence = toLowerCase(sentence)
	sentence = removePunctuation(sentence)

	if choice == 1{
		wordCount := wordCounter(sentence)

		fmt.Println("Word Count: ", wordCount)
	}else if choice == 2{
		if isPalindrome(sentence) {
			fmt.Println("The word " ,sentence, " is a palindrome.")
		} else {
			fmt.Println("The word " ,sentence, " is not palindrome.")
		}
	}

}

func removePunctuation(word string) string {
	var result string
	for i := 0; i < len(word); i++ {
		if string(word[i]) == " " || int(word[i]) >= 65 && int(word[i]) <= 90 || int(word[i]) >= 97 && int(word[i]) <= 122 {
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


func wordCounter(sentence string) map[string]int {
	var array = strings.Fields(sentence)
	var dct = make(map[string]int)

	for _, word := range array {
		dct[word]++
	}
	return dct
}

func isPalindrome(word string) bool {
	var reversedWord string

	for i := len(word) - 1; i >= 0; i-- {
		reversedWord += string(word[i])
	}

	return word == reversedWord
}
