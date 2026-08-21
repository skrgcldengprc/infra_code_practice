package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	CompanyName string
	Department  string
}

func main() {

	employee := Employee{
		Person: Person{
			Name: "テスト三郎",
			Age:  35,
		},
		CompanyName: "テスト株式会社",
		Department:  "CEO",
	}
	fmt.Println(employee.Name)
	fmt.Println(employee.Age)
	fmt.Println(employee.CompanyName)
	fmt.Println(employee.Department)
}
