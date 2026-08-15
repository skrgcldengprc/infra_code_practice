# 初心者向けのおすすめプログラム

## 1. まず作るといいもの
- 年齢判定
- ランダムな数字当てゲーム
- 電卓
- 簡単な ToDo リスト
- 単語の出現回数計算
- じゃんけん

## 2. なぜいいか
- if / for / slice / map を体験できる
- 入出力とロジックの基礎が学べる
- すぐ動くのでモチベーションが上がる

## 3. 例: 簡単な電卓
```go
package main

import "fmt"

func main() {
    a := 10
    b := 3
    fmt.Println("足し算:", a+b)
    fmt.Println("引き算:", a-b)
    fmt.Println("掛け算:", a*b)
    fmt.Println("割り算:", a/b)
}
```

## 4. 例: ToDo リスト
```go
package main

import "fmt"

func main() {
    todos := []string{"買い物", "掃除", "勉強"}
    for i, todo := range todos {
        fmt.Println(i+1, todo)
    }
}
```

## 5. まとめ
- 最初は「動くもの」を作ることが大事
- できたら次に「関数に分ける」「エラー処理を入れる」まで進める
- その繰り返しで実力がつく
