package main

import "fmt"

func main() {
	fmt.Println("問題2: 条件分岐")
	fmt.Println("以下の要件を満たすプログラムを作成してください。")
	fmt.Println("1. serverCount を定義する")
	fmt.Println("2. serverCount が 3 以上なら '運用対象が多いです' と表示する")
	fmt.Println("3. それ以外なら '運用対象は少ないです' と表示する\n")

	serverCount := 5

	if serverCount >= 3 {
		fmt.Println("運用対象が多いです")
	} else {
		fmt.Println("運用対象は少ないです")
	}

	// TODO: serverCount を使って if 文を作成してください。
	// serverCount := 5
	// if serverCount >= 3 {
	//     fmt.Println("運用対象が多いです")
	// } else {
	//     fmt.Println("運用対象は少ないです")
	// }
}
