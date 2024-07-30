package main

import (
	"bufio"
	"fmt"
	"os"
)



func main() {
	fmt.Print("Enter Your Name: ")
	var name = acceptInput()
	fmt.Print("Enter Your Number of Subjects: ")
	var numberOfSubjects float64
	fmt.Scanf("%f", &numberOfSubjects)

	var dct = make(map[string]float64)
	for i := 0; i < int(numberOfSubjects); i++ {
		var subject string
		var marks float64

		fmt.Print("Enter Subject ", i+1, " Name: ")
		fmt.Scanf("%s ", &subject)
		fmt.Print("Enter Marks: ")
		fmt.Scanf("%f", &marks)

		for {

			if marks < 0 || marks > 100 {
				fmt.Println("Enter Valid Grades that is greater than or equals 0 and less than or equals 100")
				fmt.Print("Enter Valid Mark for ", subject, ": ")
				fmt.Scanf("%f", &marks)
			} else {
				break
			}
		}

		dct[subject] = marks
	}

	average := calculateAverage(dct)
	fmt.Println("Dear ", name, " Your Total Average Score is: ", average)
}

func acceptInput() string{
	reader := bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString('\n')

	return word
}

func calculateAverage(dct map[string]float64) float64 {
	var sumOfMarks float64
	for _, value := range dct {
		sumOfMarks += value
	}
	return sumOfMarks / float64(len(dct))
}
