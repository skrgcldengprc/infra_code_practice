package main

import "fmt"

func main() {
	a := 4
	if a > 5 {
		fmt.Println("aは5より大きいです。")
	} else if a < 5 {
		fmt.Println("aは5より小さいです。")
	}
}
