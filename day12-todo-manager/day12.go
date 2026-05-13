package main

import "fmt"

type Todo struct {
	title string
	done  bool
}

func main() {
	var todos []Todo

	for {
		fmt.Println("\n--- Day 12 Todo Manager ---")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {

		case 1:
			var title string

			fmt.Print("Enter task: ")
			fmt.Scanln(&title)

			todo := Todo{
				title: title,
				done:  false,
			}

			todos = append(todos, todo)

			fmt.Println("Task added!")

		case 2:
			fmt.Println("\n--- Tasks ---")

			if len(todos) == 0 {
				fmt.Println("No tasks yet.")
			}

			for i, todo := range todos {

				status := "NOT DONE"

				if todo.done {
					status = "DONE"
				}

				fmt.Println(i+1, "-", todo.title, status)
			}

		case 3:

			fmt.Println("\n--- Tasks ---")

			for i, todo := range todos {

				status := "NOT DONE"

				if todo.done {
					status = "DONE"
				}

				fmt.Println(i+1, "-", todo.title, status)
			}

			var number int

			fmt.Print("Enter task number: ")
			fmt.Scanln(&number)

			if number > 0 && number <= len(todos) {

				todos[number-1].done = true

				fmt.Println("Task marked completed!")

			} else {
				fmt.Println("Invalid number")
			}

		case 4:

			fmt.Println("\n--- Tasks ---")

			for i, todo := range todos {

				status := "NOT DONE"

				if todo.done {
					status = "DONE"
				}

				fmt.Println(i+1, "-", todo.title, status)
			}

			var number int

			fmt.Print("Enter task number to delete: ")
			fmt.Scanln(&number)

			if number > 0 && number <= len(todos) {

				todos = append(todos[:number-1], todos[number:]...)

				fmt.Println("Task deleted!")

			} else {
				fmt.Println("Invalid number")
			}

		case 5:
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}
