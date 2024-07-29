package main

import "fmt"


func main(){
	var revenue float64
	var expense float64
	const taxRate float64 = 0.5

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter Expense: ")
	fmt.Scan(&expense)

	var ebt = revenue - expense
	var profit = ebt * (1-taxRate)


	fmt.Println("Profit is",profit)
	fmt.Println("EBT is",ebt)
}