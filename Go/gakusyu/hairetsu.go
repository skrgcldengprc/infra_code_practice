package main

//複数のデータをまとめて扱いたい場合、配列を利用すると効率的な処理が書ける。

import ("fmt")

func main() {
	// 配列の宣言
	// var scores [5]int
	// scores[0] = 10
	// scores[1] = 20
	// scores[2] = 30
	// scores[3] = 40
	// scores[4] = 50

	// 配列の要素を出力
	// for i := 0; i < len(scores); i++ {
	// 	fmt.Println(scores[i])

		
    //配列の初期値は、定義する際にまとめて設定することができる。
	//scores := [5]int{10, 20, 30, 40, 50}
	
	// for i := 0; i < len(scores); i++ {
	// 	fmt.Println(scores[i])
	//     }

	//配列で特殊な繰り返しを利用できる
    scores := [4]int{10,20,30,99}

	for i, score := range scores {
		fmt.Println(i)
		fmt.Println(score)
	}
}


