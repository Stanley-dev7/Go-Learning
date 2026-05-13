package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	title string
	done  bool
}

func saveTodos(todos []Todo) {

	file, _ := os.Create("todos.txt")
	defer file.Close()

	for _, todo := range todos {

		line := todo.title + "," + strconv.FormatBool(todo.done) + "\n"

		file.WriteString(line)
	}
}

func loadTodos() []Todo {

	var todos []Todo

	file, err := os.Open("todos.txt")

	if err != nil {
		return todos
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		parts := strings.Split(line, ",")

		if len(parts) == 2 {

			done, _ := strconv.ParseBool(parts[1])

			todo := Todo{
				title: parts[0],
				done:  done,
			}

			todos = append(todos, todo)
		}
	}

	return todos
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	todos := loadTodos()

	for {

		fmt.Println("\n--- Day 13 Persistent Todo Manager ---")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Mark Task Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")

		fmt.Print("Enter choice: ")

		choiceInput, _ := reader.ReadString('\n')
		choiceInput = strings.TrimSpace(choiceInput)

		choice, err := strconv.Atoi(choiceInput)

		if err != nil {
			fmt.Println("Please enter a number")
			continue
		}

		switch choice {

		case 1:

			fmt.Print("Enter task: ")

			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			todo := Todo{
				title: title,
				done:  false,
			}

			todos = append(todos, todo)

			saveTodos(todos)

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

			fmt.Print("Enter task number: ")

			numberInput, _ := reader.ReadString('\n')
			numberInput = strings.TrimSpace(numberInput)

			number, err := strconv.Atoi(numberInput)

			if err != nil {
				fmt.Println("Please enter a valid number")
				continue
			}

			if number > 0 && number <= len(todos) {

				todos[number-1].done = true

				saveTodos(todos)

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

			fmt.Print("Enter task number to delete: ")

			numberInput, _ := reader.ReadString('\n')
			numberInput = strings.TrimSpace(numberInput)

			number, err := strconv.Atoi(numberInput)

			if err != nil {
				fmt.Println("Please enter a valid number")
				continue
			}

			if number > 0 && number <= len(todos) {

				todos = append(todos[:number-1], todos[number:]...)

				saveTodos(todos)

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
