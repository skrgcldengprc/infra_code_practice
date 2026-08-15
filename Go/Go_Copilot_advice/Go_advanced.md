# Go上級者編

## 1. 目標
- 並行処理を安全に扱える
- API や DB を使える
- 設計と保守性を意識できる
- パッケージ構成を整理できる

## 2. 学ぶべきテーマ
- goroutine と channel
- mutex / RWMutex
- context
- net/http
- database/sql
- testing
- benchmarking
- logging / config / env

## 3. 学習順
1. goroutine と channel
2. sync パッケージ
3. context
4. HTTP サーバー
5. DB 接続
6. テストコードを書く
7. エラー設計とログ設計

## 4. 実務で作ると良い例
- REST API
- バッチ処理
- 社内ツール
- ログ収集CLI
- API クライアント

### 例
```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "hello")
}

func main() {
    http.HandleFunc("/", handler)
    _ = http.ListenAndServe(":8080", nil)
}
```
