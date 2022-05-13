package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

		if strings.Compare(userInput, "5") == 0 {
			fmt.Println("Bye!")
			os.Exit(0)
		}
	}

}
