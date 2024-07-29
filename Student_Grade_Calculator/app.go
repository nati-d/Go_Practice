package main

import "fmt"

func main() {
	fmt.Print("Enter Your Name: ")
	var name string
	fmt.Scanf("%s", &name)
	fmt.Print("Enter Your Number of Subjects: ")
	var numberOfSubjects float64
	fmt.Scanf("%f", &numberOfSubjects)

	var dct = make(map[string]float64)
	for i := 0; i < int(numberOfSubjects); i++ {
		var subject string
		var marks float64

		fmt.Print("Enter Subject Name: ")
		fmt.Scanf("%s ", &subject)
		fmt.Print("Enter Marks: ")
		fmt.Scanf("%f", &marks)
		dct[subject] = marks
	}

	average := calculateAverage(dct)
	fmt.Println("Dear ", name, " Your Total Average Score is: ", average)
}

func calculateAverage(dct map[string]float64) float64 {
	var sumOfMarks float64
	for _, value := range dct {
		sumOfMarks += value
	}
	return sumOfMarks / float64(len(dct))
}
