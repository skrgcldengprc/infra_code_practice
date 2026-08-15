# Go中級者編

## 1. 目標
- 関数を分けて整理できる
- error をちゃんと返せる
- file, JSON, HTTP の基礎が触れる
- package を分割して扱える

## 2. 学ぶべきポイント
- 関数分割と責務分離
- struct の設計
- slice / map の応用
- error handling
- defer / panic / recover の使いどころ
- JSON の読み書き

## 3. 典型的な学習順
1. 関数を複数ファイルへ分割
2. struct を使ってデータ管理
3. error を返す関数を作る
4. JSON を扱う
5. ファイル入出力
6. 標準ライブラリの活用

## 4. 役立つ小さなアプリ
- ToDoアプリ
- メモ帳
- CSV 読み取り
- 簡単な電卓
- 単語カウント

### 例
```go
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main() {
    u := User{Name: "Aoi", Age: 25}
    b, _ := json.Marshal(u)
    fmt.Println(string(b))
}
```
