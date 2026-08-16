package main

import (
	"fmt"
)

func main() {
	fmt.Println("問題8: ディレクトリ一覧表示")
	fmt.Println("1. 現在のディレクトリのファイル一覧を取得する")
	fmt.Println("2. それを1つずつ表示する")

	// TODO: os.ReadDir を使って一覧を表示してください。
	// entries, err := os.ReadDir(".")
	// if err != nil {
	//     fmt.Println("読み込みエラー:", err)
	//     return
	// }
	// for _, entry := range entries {
	//     fmt.Println(entry.Name())
	// }
}
