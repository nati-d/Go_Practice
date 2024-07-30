package main

import "fmt"

func main() {
	for {
	fmt.Println("\n What Do You Want To Do? ")
	fmt.Println("1. Count Words")
	fmt.Println("2. Check Palindrome")
	var choice int
	fmt.Print("Enter Your Choice: ")
	fmt.Scanln(&choice)

	if choice == 1 {
		wordCounterFunction()
	} else if choice == 2 {
		palindromChecker()
	}
}

}
