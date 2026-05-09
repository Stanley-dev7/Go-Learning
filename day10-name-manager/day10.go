package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file := "names.txt"
	// load existing names
	data, _ := os.ReadFile(file)
	names := strings.Split(string(data), "\n")
	for {
		fmt.Println("\n--- Day 10 Manager---")
		fmt.Println("1. Add name")
		fmt.Println("2. view names")
		fmt.Println("3. Exist")
		fmt.Print("Enter choice:")
		var choice int
		fmt.Scanln(&choice)
		switch choice {

		case 1:
			var name string
			fmt.Print("Enter name:")
			fmt.Scanln(&name)
			names = append(names, name)
			// save to file
			os.WriteFile(file, []byte(strings.Join(names, "\n")), 0644)
			fmt.Println("Name saved!")
		case 2:
			fmt.Println("\n--- Names list---")
			for i, name := range names {
				if name != "" {
					fmt.Println(i+1, name)
				}
			}
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Ivalid option")
		}
	}
}
