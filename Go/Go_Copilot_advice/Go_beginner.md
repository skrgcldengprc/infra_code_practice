# Go入門編

## 1. ここまでを目標にする
- 変数、定数、if、for、関数が書ける
- slice、map、struct を理解する
- error を正しく扱える
- go run と gofmt を使える

## 2. 学習順
1. package / import / func main
2. 変数と型
3. if / for
4. 関数
5. slice / map / struct
6. error
7. pointer
8. package とファイル分割

## 3. なぜ重要か
- Goは型が明確で、ミスが見つけやすい
- シンプルな文法が多い
- 実務でよく使う構成はこの流れで学ぶ

## 4. まず作るといいプログラム
- 年齢判定
- 料金計算
- じゃんけん
- 素数判定
- 配列の合計

### 例
```go
package main

import "fmt"

func main() {
    nums := []int{10, 20, 30}
    sum := 0
    for _, n := range nums {
        sum += n
    }
    fmt.Println(sum)
}
```
