package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("問題10: サーバー状態の簡単なチェック")
	fmt.Println("1. URL を指定して HTTP リクエストを送る")
	fmt.Println("2. 200 OK なら '正常' と表示する")
	fmt.Println("3. それ以外なら '異常' と表示する")

	// TODO: http.Get を使って HTTP サーバーの状態を確認してください。
	// resp, err := http.Get("https://example.com")
	// if err != nil {
	//     fmt.Println("通信エラー:", err)
	//     return
	// }
	// defer resp.Body.Close()
	// if resp.StatusCode == http.StatusOK {
	//     fmt.Println("正常")
	// } else {
	//     fmt.Println("異常")
	// }
	_ = http.StatusOK
}
