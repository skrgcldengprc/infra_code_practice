package main

import "fmt"

func main() {
	// 変数の宣言
	var a int = 10
	var b int = 20
	// 変数の値を出力
	fmt.Println("aの値は", a)
	fmt.Println("bの値は", b)
	fmt.Println("a + bの値は", a+b)
}

// 変数を宣言して同時に値を初期化するときは:=を使う。
func main() {
	// 変数の宣言と初期化
	a := 10
}

//変数を宣言するだけで初期化しない場合は、varを使う。
func main() {
	// 変数の宣言
	var a int
}

//変数を初期化する値がわからないときに先に宣言する必要があるときは、varを使う。
func main() {
	// 変数の宣言
	var a custom = 5
	fmt.Println(a)
}

//ファイルのトップでmain関数外の時には、var()を使って複数宣言することが多い。

var (
	a string
	b string
	c string
)

func main() {
	a = "Hello"
	b = "World"
	c = "!"
}
