package main

import "fmt"

func main() {
	var choice int
	var amount int
	balance := 1000
	for {
		fmt.Println("\n--- Atm menu---")
		fmt.Println("1. check balance")
		fmt.Println("2. deposit")
		fmt.Println("3.withdraw")
		fmt.Println("4.exist")
		fmt.Println("Enter choice:")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is:", balance)
		case 2:
			fmt.Println("Enter deposit amount:")
			fmt.Scanln(&amount)
			balance += amount
			fmt.Println("Deposited succesfully. New balance", balance)
		case 3:
			fmt.Println("Enter withdawal amount:")
			fmt.Scanln(&amount)
			if amount <= balance {
				balance -= amount
				fmt.Println("Withdrwal successful. New balance", balance)
			} else {
				fmt.Println("Insufficient balance")

			}

		case 4:
			fmt.Println("goodbye")
			return
		default:
			fmt.Println("Invalid option")
		}

	}
}
