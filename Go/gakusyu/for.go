package main

import "fmt"

func main() {
	index := 0

	for index < 10 {
		if index%2 == 0 {
			fmt.Println(index, "valueは偶数です。")
		} else if index%2 != 0 {
			fmt.Println(index)
		} else {
			fmt.Println("範囲外です。")
		}
		index = index + 1
	}
	//効率の良い書き方

	for index := 0; index < 10; index++ {
		if index%2 == 0 {
			fmt.Println(index, "valueは偶数です。")
			continue
		} else if index%2 != 0 {
			fmt.Println(index)
		} else {
			fmt.Println("範囲外です。")
		}
		fmt.Println(index, "continue走ってるから奇数をもう1回表示")
	}
}

//break文を使うと、ループを途中で抜けることができる。
//continue文を使うと、繰り返しの途中で次のループに移る。
