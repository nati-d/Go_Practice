package main

import "math"
import "fmt"

func main(){
	investmentAmount,years  := 1000,10
	var expectdReturnRate = 5.5
	// var years = 10

	var futureValue = float64(investmentAmount) * math.Pow(1 + expectdReturnRate / 100,float64(years))

	fmt.Println("Future value is",int(futureValue))
}