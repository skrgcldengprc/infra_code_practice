package main

import "fmt"

func main() {
	// int型の変数を宣言(整数を格納するための型)
	var a int = 10
	// float64型の変数を宣言(froat32型は小数点以下の精度が低いため、float64型を使用することが一般的)
	var b float64 = 20.5
	// string型の変数を宣言(文字列を格納するための型)
	var c string = "Hello"
	// bool型の変数を宣言(trueまたはfalseの値を持つ)
	var d bool = true

	//型推論
	e := 30 // 型推論による変数宣言

	// 変数の値を出力
	fmt.Println("aの値は", a)
	fmt.Println("bの値は", b)
	fmt.Println("cの値は", c)
	fmt.Println("dの値は", d)
	fmt.Println("eの値は", e)
}
