package main

import (
	"fmt"
	"go-example/patterns/design-patterns/behavioural/visitor"
)

func main() {
	employees := []visitor.Employee{
		&visitor.Manager{Name: "John", Salary: 5000},
		&visitor.Engineer{Name: "Mary", Salary: 4000},
		&visitor.Salesman{Name: "Bob", Salary: 3000},
	}

	calculator := &visitor.SalaryCalculator{}

	for _, employee := range employees {
		employee.Accept(calculator)
	}

	fmt.Println("Total salary:", calculator.TotalSalary)
}
