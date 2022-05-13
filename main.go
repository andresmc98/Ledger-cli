package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Welcome to Ledger CLI")
	fmt.Println("------------------")
	fmt.Println("1. Make a new income")
	fmt.Println("2. Make a new expense")
	fmt.Println("3. Show all incomes")
	fmt.Println("4. Show all expenses")
	fmt.Println("5. exit")
	fmt.Println("=>")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		userInput := scanner.Text()
		fmt.Println(userInput)
		fmt.Print("=>")

		switch userInput {
		case "1":
			fmt.Println("You have selected Make a new income")
		case "2":
			fmt.Println("You have selected Make a new expense")
		case "3":
			fmt.Println("You have selected Show all incomes")
			read("./fixtures/Income.ledger")
		case "4":
			fmt.Println("You have selected Show all expenses")
			read("fixtures/Expenses.ledger")
		case "5":
			fmt.Println("You have selected exit")
			fmt.Println("Bye!")
			os.Exit(0)
		}
	}

}
