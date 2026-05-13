package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Todo struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func saveTodos(todos []Todo) {

	data, err := json.MarshalIndent(todos, "", "  ")

	if err != nil {
		fmt.Println("Error saving todos")
		return
	}

	os.WriteFile("todos.json", data, 0644)
}

func loadTodos() []Todo {

	var todos []Todo

	data, err := os.ReadFile("todos.json")

	if err != nil {
		return todos
	}

	json.Unmarshal(data, &todos)

	return todos
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	todos := loadTodos()

	for {

		fmt.Println("\n--- Day 14 JSON Todo Manager ---")
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
			fmt.Println("Please enter a valid number")
			continue
		}

		switch choice {

		case 1:

			fmt.Print("Enter task: ")

			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			todo := Todo{
				Title: title,
				Done:  false,
			}

			todos = append(todos, todo)

			saveTodos(todos)

			fmt.Println("Task added!")

		case 2:

			fmt.Println("\n--- Tasks ---")

			if len(todos) == 0 {
				fmt.Println("No tasks found.")
			}

			for i, todo := range todos {

				status := "❌"

				if todo.Done {
					status = "✅"
				}

				fmt.Println(i+1, "-", todo.Title, status)
			}

		case 3:

			fmt.Println("\n--- Tasks ---")

			for i, todo := range todos {

				status := "❌"

				if todo.Done {
					status = "✅"
				}

				fmt.Println(i+1, "-", todo.Title, status)
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

				todos[number-1].Done = true

				saveTodos(todos)

				fmt.Println("Task completed!")

			} else {
				fmt.Println("Invalid number")
			}

		case 4:

			fmt.Println("\n--- Tasks ---")

			for i, todo := range todos {

				status := "❌"

				if todo.Done {
					status = "✅"
				}

				fmt.Println(i+1, "-", todo.Title, status)
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
