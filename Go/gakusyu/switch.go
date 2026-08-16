package main

import "fmt"

func main() {
	value := 3
	
	switch value {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Other")
	}
}