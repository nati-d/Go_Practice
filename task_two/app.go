package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
	"github.com/fatih/color"
)

func main() {
	for {
		mainMenu()
		var continueChoice string
		fmt.Print(color.CyanString("\nDo you want to continue? (y/n): "))
		fmt.Scanln(&continueChoice)

		if continueChoice == "n" {
			fmt.Println(color.RedString("Thank you for using the application!"))
			break
		} else if continueChoice != "y" {
			fmt.Println(color.YellowString("Invalid choice. Please enter 'y' for yes or 'n' for no."))
		}

		clearScreen()
	}
}

func mainMenu() {
	var boldCyan = color.New(color.FgCyan, color.Bold).SprintFunc()
	var boldMagenta = color.New(color.FgMagenta, color.Bold).SprintFunc()

	fmt.Println(boldCyan("Welcome to the Word Counter and Palindrome Checker Application!"))
	fmt.Println(color.BlueString("What do you want to do?"))
	fmt.Println("1. Count Words")
	fmt.Println("2. Check Palindrome")
	fmt.Println("3. Exit")
	var choice int
	fmt.Print(boldMagenta("Enter Your Choice: "))
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		wordCounterFunction()
	case 2:
		palindromChecker()
	case 3:
		fmt.Println(color.YellowString("Exiting the application..."))
		os.Exit(0)
	default:
		fmt.Println(color.RedString("Invalid choice. Please enter the Correct Choice."))
	}
}

func clearScreen() {
	time.Sleep(500 * time.Millisecond) 
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}


