package main


import "fmt"

func main () {
	var accountBalance float64 = 1000
	outputText("Welcome to the Bank!")
	outputText("Please select an option:")
	outputText("1. Check Balance")
	outputText("2. Deposit")
	outputText("3. Withdraw")
	outputText("4. Exit")



	var choice int 
	fmt.Scan(&choice)

	fmt.Println("You selected",choice)

	if choice == 1 {
		fmt.Println("Your account balance is",accountBalance)
	}else if choice == 2{
		var depositAmount float64
		fmt.Print("Enter deposit amount: ")
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount
		fmt.Printf("Your new balance is %.1f",accountBalance)
	}else if choice == 3{
		var withdrawAmount float64
		fmt.Print("Enter withdraw amount: ")
		fmt.Scan(&withdrawAmount)
		if withdrawAmount > accountBalance {
			fmt.Println("Insufficient balance")
		}else{
			accountBalance -= withdrawAmount
			fmt.Println("Your new balance is",accountBalance)
		}
	} else{
		fmt.Println("Goodbye!")
	}

}

func outputText(text string){
	fmt.Println(text)
}