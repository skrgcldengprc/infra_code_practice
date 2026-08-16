package main

import (
	"fmt"
)

func main() {
	fmt.Println("問題7: ファイル読み込み")
	fmt.Println("1. 'sample.txt' というファイルを読み込む")
	fmt.Println("2. 内容を1行ずつ表示する")
	fmt.Println("3. ファイルが存在しない場合はエラーを表示する")

	// TODO: os.ReadFile か os.Open を使って内容を表示してください。
	// data, err := os.ReadFile("sample.txt")
	// if err != nil {
	//     fmt.Println("読み込みエラー:", err)
	//     return
	// }
	// fmt.Println(string(data))
}
