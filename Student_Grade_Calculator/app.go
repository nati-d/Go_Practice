package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"os"
	"strings"
)

func main() {
	fmt.Println("------------------------------------------------------------")
	fmt.Println("          Welcome to the Grade Calculator!                  ")
	fmt.Println("------------------------------------------------------------")

	fmt.Print("Please enter your name: ")
	name := strings.TrimSpace(acceptInput())

	var numberOfSubjects int
	for {
		fmt.Print("How many subjects do you have, ", name, "? ")
		fmt.Scanf("%d", &numberOfSubjects)
		if numberOfSubjects > 0 {
			break
		} else {
			fmt.Println("Please enter a valid number of subjects (greater than 0).")
		}
	}

	dct := make(map[string]float64)
	for i := 0; i < numberOfSubjects; i++ {
		var subject string
		var marks float64

		fmt.Printf("Enter the name of subject %d: ", i+1)
		fmt.Scanf("%s", &subject)

		for {
			fmt.Printf("Enter your marks for %s (0-100): ", subject)
			fmt.Scanf("%f", &marks)
			if marks >= 0 && marks <= 100 {
				break
			} else {
				fmt.Println("Invalid marks! Please enter a value between 0 and 100.")
			}
		}

		dct[subject] = marks
	}
	buildTable(dct, name)
	fmt.Println("------------------------------------------------------------")
	var message string
	if calculateAverage(dct) >= 50 {
		message = "Congratulations! You have passed."
	} else {
		message = "Sorry! You have failed."
	}
	fmt.Printf("Your average grade is %.2f\n", calculateAverage(dct))
	fmt.Println(message)

}

func acceptInput() string {
	reader := bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString('\n')
	return strings.TrimSpace(word)
}

func calculateAverage(dct map[string]float64) float64 {
	var sumOfMarks float64
	for _, value := range dct {
		sumOfMarks += value
	}
	return sumOfMarks / float64(len(dct))
}

func buildTable(dct map[string]float64, name string) {
	average := calculateAverage(dct)

	fmt.Printf("\n%s's Grade Report\n", name)
	tbl := table.New("ID", "Subject", "Mark")
	id := 1
	for key, value := range dct {
		headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
		columnFmt := color.New(color.FgYellow).SprintfFunc()
		tbl.AddRow(id, key, value)
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt).WithPadding(6)
		id++
	}
	tbl.AddRow("*", "Average", average)
	tbl.Print()
}
