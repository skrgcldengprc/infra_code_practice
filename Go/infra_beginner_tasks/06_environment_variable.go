package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("問題6: 環境変数")
	fmt.Println("1. 環境変数 'HOME' または 'USER' を取得する")
	fmt.Println("2. その値を表示する")
	fmt.Println("3. なければ '環境変数が見つかりません' と表示する")

	value, ok :=os.LookupEnv("HOME")
	if ok{
		fmt.Println("HOME =", value)
	}else {
		fmt.Println("環境変数が見つかりません")
	}

	// TODO: os.Getenv を使って環境変数を取得してください。
	// value, ok := os.LookupEnv("HOME")
	// if ok {
	//     fmt.Println("HOME =", value)
	// } else {
	//     fmt.Println("環境変数が見つかりません")
	// }
}
