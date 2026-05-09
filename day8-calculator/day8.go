package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func subtract(a int, b int) int {
	return a - b
}
func multiply(a int, b int) int {
	return a * b
}
func divide(a int, b int) int {
	return a / b
}
func main() {
	var choice int
	var num1, num2 int
	for {
		fmt.Println("/n---calculator---")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exist")
		fmt.Println("Enter your choice:")
		fmt.Scanln(&choice)
		if choice == 5 {
			fmt.Println("Goodbye")
			return
		}
		fmt.Println("Enter first number:")
		fmt.Scanln(&num1)
		fmt.Println("Enter second number:")
		fmt.Scanln(&num2)
		switch choice {
		case 1:
			fmt.Println("Result:", add(num1, num2))
		case 2:
			fmt.Println("Result:", subtract(num1, num2))
		case 3:
			fmt.Println("Result:", multiply(num1, num2))
		case 4:
			if num2 != 0 {
				fmt.Println("Result", divide(num1, num2))
			} else {
				fmt.Println("cannot divide by 0")
			}
		default:
			fmt.Println("Invalid option")
		}
	}

}
