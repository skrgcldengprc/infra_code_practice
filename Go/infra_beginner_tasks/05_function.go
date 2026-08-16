package main

import "fmt"

func main() {
	fmt.Println("問題5: 関数")
	fmt.Println("1. addNumbers という関数を作る")
	fmt.Println("2. 2つの整数を受け取り、足し算の結果を返す")
	fmt.Println("3. main 関数から呼び出して表示する")

	// TODO: addNumbers を定義し、呼び出してください。
	// result := addNumbers(10, 25)
	// fmt.Println("結果:", result)
	result := addNumbers(10, 25)
	fmt.Println("結果:", result)
}

func addNumbers(value1 int, value2 int) int {
	return value1 + value2
}

// func addNumbers(a int, b int) int {
//     return a + b
// }
