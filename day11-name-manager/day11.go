package main

import (
	"fmt"
	"os"
	"strings"
)

func saveNames(file string, names []string) {
	os.WriteFile(file, []byte(strings.Join(names, "\n")), 0644)
}

func main() {

	file := "names.txt"

	data, _ := os.ReadFile(file)

	content := strings.TrimSpace(string(data))

	var names []string

	if content != "" {
		names = strings.Split(content, "\n")
	}

	for {

		fmt.Println("\n--- Day 11 Name Manager ---")
		fmt.Println("1. Add Name")
		fmt.Println("2. View Names")
		fmt.Println("3. Delete Name")
		fmt.Println("4. Search Name")
		fmt.Println("5. Edit Name")
		fmt.Println("6. Exit")

		fmt.Print("Enter choice: ")

		var choice int

		_, err := fmt.Scan(&choice)

		if err != nil {
			fmt.Println("Please enter a number")

			var clear string
			fmt.Scan(&clear)

			continue
		}

		switch choice {

		case 1:

			var name string

			fmt.Print("Enter name: ")
			fmt.Scan(&name)

			names = append(names, name)

			saveNames(file, names)

			fmt.Println("Name added!")

		case 2:

			if len(names) == 0 {

				fmt.Println("No names found")

			} else {

				fmt.Println("\n--- Names List ---")

				for i, n := range names {
					fmt.Println(i+1, n)
				}
			}

		case 3:

			if len(names) == 0 {

				fmt.Println("No names to delete")
				continue
			}

			fmt.Println("\n--- Names List ---")

			for i, n := range names {
				fmt.Println(i+1, n)
			}

			var index int

			fmt.Print("Enter number to delete: ")
			fmt.Scan(&index)

			if index >= 1 && index <= len(names) {

				names = append(names[:index-1], names[index:]...)

				saveNames(file, names)

				fmt.Println("Name deleted!")

			} else {

				fmt.Println("Invalid number")
			}

		case 4:

			var search string

			found := false

			fmt.Print("Enter name to search: ")
			fmt.Scan(&search)

			for i, n := range names {

				if strings.EqualFold(n, search) {

					fmt.Println("Found at position:", i+1)

					found = true
				}
			}

			if !found {
				fmt.Println("Name not found")
			}

		case 5:

			if len(names) == 0 {

				fmt.Println("No names to edit")

				continue
			}

			fmt.Println("\n--- Names List ---")

			for i, n := range names {
				fmt.Println(i+1, n)
			}

			var index int

			fmt.Print("Enter number to edit: ")
			fmt.Scan(&index)

			if index >= 1 && index <= len(names) {

				var newName string

				fmt.Print("Enter new name: ")
				fmt.Scan(&newName)

				names[index-1] = newName

				saveNames(file, names)

				fmt.Println("Name updated!")

			} else {

				fmt.Println("Invalid number")
			}

		case 6:

			fmt.Println("Goodbye!")

			return

		default:

			fmt.Println("Invalid option")
		}
	}
}
