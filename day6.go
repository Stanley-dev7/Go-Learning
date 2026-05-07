package main

import "fmt"

func main() {
	employee := []string{"Stanley", "John"}
	employee = append(employee, "Mary", "David")
	fmt.Println("Employee list:")
	for index, employee := range employee {
		fmt.Println(index+1, employee)
	}

}
