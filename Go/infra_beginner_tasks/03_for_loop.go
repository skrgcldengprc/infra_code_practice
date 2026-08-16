package main

import "fmt"

func main() {
	fmt.Println("問題3: 繰り返し処理")
	fmt.Println("1. 1 から 10 までを順番に表示してください")
	fmt.Println("2. その後、合計値を表示してください")

	total := 0

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		total += i
	}

	fmt.Println("合計:", total)

	// TODO: for ループで 1 から 10 まで出力し、合計を計算してください。
	// total := 0
	// for i := 1; i <= 10; i++ {
	//     fmt.Println(i)
	//     total += i
	// }
	// fmt.Println("合計:", total)
}
