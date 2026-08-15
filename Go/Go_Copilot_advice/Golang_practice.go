package main

import (
	"errors"
	"fmt"
	"os"
)

// Go基礎まとめ
// 何度も見返すための簡潔メモ

// 1. 基本構文
// - package main: 実行用のパッケージ
// - import: 外部パッケージの読み込み
// - func main(): 実行の入口
// - Goはセミコロンを省略できる

func main() {
	fmt.Println("Go基礎まとめ")

	// 2. 変数
	// var で宣言 or := で短く宣言
	var age int = 20
	name := "Alice"
	fmt.Println(age, name)

	// 3. 条件分岐
	if age >= 20 {
		fmt.Println("成人です")
	} else {
		fmt.Println("未成年です")
	}

	// 4. ループ
	// Goの for はC言語風
	for i := 0; i < 3; i++ {
		fmt.Println("for:", i)
	}

	// 5. 関数
	fmt.Println("add:", add(3, 4))

	// 6. 配列とスライス
	// 配列は長さ固定、スライスは可変長
	arr := [3]int{1, 2, 3}
	nums := []int{10, 20, 30}
	nums = append(nums, 40)
	fmt.Println(arr, nums)

	// 7. マップ
	score := map[string]int{"A": 90, "B": 80}
	fmt.Println("score:", score["A"])

	// 8. 構造体
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Bob", Age: 30}
	fmt.Println("person:", p.Name, p.Age)

	// 9. ポインタ
	// & はアドレス取得、* は参照
	p2 := &p
	p2.Age = 31
	fmt.Println("pointer:", p.Age)

	// 10. エラー処理
	// Goでは例外ではなく error を返すのが基本
	_, err := os.Open("no-file")
	if err != nil {
		fmt.Println("error:", err)
	}

	// 11. defer
	// 関数終了時に最後に実行される
	defer fmt.Println("defer: 関数終了前に実行")

	// 12. goroutine
	// 並行処理の基本。軽いスレッドのようなもの
	ch := make(chan string, 1)
	go func() {
		ch <- "goroutine OK"
	}()
	fmt.Println(<-ch)

	// 13. 補足
	// - Goは静的型付け
	// - 変数は基本的に値渡し
	// - スライス・マップは扱いに注意
	// - fmt.Println で簡単に出力できる
	// - gofmt で整形するのが基本
	// - go run: 実行, go build: ビルド, go test: テスト
	fmt.Println("補足: Goはシンプルだが、設計思想が大事")
}

// 5. 関数の例
func add(a int, b int) int {
	return a + b
}

// 10. エラーの例
func readFile(path string) (string, error) {
	if path == "" {
		return "", errors.New("empty path")
	}
	return "ok", nil
}

// まとめ
// - Goはシンプルで読みやすいが、型・エラー処理・並行処理の考え方を覚えると強い
// - 初心者は「変数→条件→関数→スライス/マップ→エラー→並行処理」の順で学ぶとわかりやすい
// - まずは go run と gofmt を使い続けることが上達の近道
