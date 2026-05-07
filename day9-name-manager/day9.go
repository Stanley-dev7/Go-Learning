package main

import "fmt"

func main() {
	var names []string
	var choice int
	var name string
	var index int

	for {
		fmt.Println("\n---- Name of managers ---")
		fmt.Println("1. Add names")
		fmt.Println("2. view names")
		fmt.Println("3.Delete names")
		fmt.Println("4. Search names")
		fmt.Println("5.Exist")
		fmt.Println("Enter chioce:")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			fmt.Println("Enter name:")
			fmt.Scanln(&name)
			names = append(names, name)
			fmt.Println("name added!")
		case 2:
			fmt.Println("List of names")
			if len(names) == 0 {
				fmt.Println("No names yet:")
			} else {
				for index, value := range names {
					fmt.Println(index+1, value)
				}

			}
		case 3:
			if len(names) == 0 {
				fmt.Println("No names to delete")
				continue
			}
			fmt.Println("Enter number to delete:")
			fmt.Scanln(&index)
			if index > 0 && index <= len(names) {
				names = append(names[:index-1], names[index:]...)
				fmt.Println("Name delected!")
			} else {
				fmt.Println("Invalid number")
			}
		case 4:
			if len(names) == 0 {
				fmt.Println("No names to search")
				continue
			}
			fmt.Println("Enter name to search:")
			fmt.Scanln(&name)
			found := false
			for i, v := range names {
				if v == name {
					fmt.Println("Found at position:", i+1)
					found = true
				}
			}
			if !found {
				fmt.Println("Name not found")
			}
		case 5:
			fmt.Println("Goodbye")
			return
		default:
			fmt.Println("Invalid option")

		}
	}
}
