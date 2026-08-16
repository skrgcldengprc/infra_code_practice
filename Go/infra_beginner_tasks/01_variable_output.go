package main

import "fmt"

func main() {
	fmt.Println("問題1: 変数と出力")
	fmt.Println("以下の要件を満たすプログラムを作成してください。")
	fmt.Println("1. name という変数を定義する")
	fmt.Println("2. age という変数を定義する")
	fmt.Println("3. 自分の名前と年齢を1行で表示する")
	fmt.Println("4. 例: 私の名前は山田太郎で、30歳です。\n")

	name := "山田太郎"
	age := 30

	fmt.Printf("私の名前は%sで、%d歳です。\n", name, age)

	// TODO: name と age を定義して、1行の文章として出力してください。
	// 例:
	// name := "山田太郎"
	// age := 30
	// fmt.Println("私の名前は" + name + "で、" + strconv.Itoa(age) + "歳です。")
}
