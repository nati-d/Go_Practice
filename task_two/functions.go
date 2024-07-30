package main

import (
	"strings"
)

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
