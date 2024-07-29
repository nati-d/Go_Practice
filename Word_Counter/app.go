package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sentence: ")
	sentence, _ := reader.ReadString('\n')

	sentence = trimWord(sentence)
	sentence = toLowerCase(sentence)
	sentence = removePunctuation(sentence)

	wordCount := wordCounter(sentence)


	fmt.Println("Word Count: ", wordCount)

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


func wordCounter(sentence string) map[string]int {
	var array = strings.Fields(sentence)
	var dct = make(map[string]int)

	for _, word := range array {
		dct[word]++
	}
	return dct
}
