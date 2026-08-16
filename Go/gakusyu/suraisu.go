package main

//要素数を追加したい場合
//配列は要素数を途中で追加できないため、要素数が不明確な場合はスライスを利用する。
//スライスを使うと要素数は固定しない

import "fmt"

func main() {
	//var scores []int
	//まとめて宣言するには
	scores := []int{10, 20, 30, 40, 50}

	scores = append(scores, 80)
	scores = append(scores, 99)

	// scores = append(scores, 10)
	// scores = append(scores, 20)
	// scores = append(scores, 30)
	// scores = append(scores, 500)

	//インデックスの変数を使わないようにするには_を使う。
	for _, score := range scores {
		//fmt.Println(i)
		fmt.Println(score)
	}
}
