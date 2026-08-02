package main

import "fmt"

func main() {
	a := 5
	if a > 5 {
		fmt.Println("aは5より大きいです。")
	} else if a < 5 {
		fmt.Println("aは5より小さいです。")
	} else {
		fmt.Println("aは5と等しいです。")
	}
}
