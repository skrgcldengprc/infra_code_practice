# Kubernetesで使うGo言語の要点まとめ

## 1. この資料の前提
- Goはバックエンド言語としても優秀だが、Kubernetesでは「controller」「client-go」「operator」「kubectl plugin」などの開発で頻繁に使われる
- Kubernetes は Go で多くのコンポーネントが実装されている
- 学習の目的は「KubernetesのAPIやリソースを触る」「自前のcontrollerやoperatorを作る」ことが中心

## 2. GoでKubernetesを学ぶときに特に重要な考え方

### (1) 構造体とJSONの理解
Kubernetes API は JSON/YAML をベースにしているため、Goの構造体とJSONタグは非常に重要。

```go
type Pod struct {
    Metadata Metadata `json:"metadata"`
    Spec     PodSpec  `json:"spec"`
}
```

- `json:"name"` のようなタグを使える
- KubernetesのAPIオブジェクトはGoのstructとして表現される
- これがわかると `client-go` を使いやすくなる

### (2) interface と振る舞いの扱い
Kubernetesでは interface を多用する構成が多い。

```go
type Client interface {
    Get() error
    Create() error
}
```

- 依存関係を抽象化しやすい
- controller や informer の設計で頻出
- 「継承」より「振る舞い」で設計する感覚が重要

### (3) error を正しく扱う
Kubernetesでは API呼び出しの失敗が普通に起こる。

```go
pod, err := clientset.CoreV1().Pods(ns).Get(ctx, name, metav1.GetOptions{})
if err != nil {
    return err
}
```

- Goのエラー処理は Kubernetes では必須
- API 失敗、認証失敗、NotFound などを分けて扱う
- `errors.Is` や `apierrors.IsNotFound` を使うと実務向き

### (4) context を理解する
Kubernetes API 操作では context が非常に重要。

```go
ctx := context.Background()
```

- timeout や cancel を制御できる
- Kubernetes client はほぼ context を受け取る
- controller 稼働やリクエスト制御で必須

## 3. Kubernetesで特に多用されるGoの文法・概念

### 1. struct
```go
type Deployment struct {
    Name string
    Replicas int32
}
```
- Kubernetesオブジェクトはほぼ struct で表現される
- YAML を Go の構造に落とし込む感覚が重要

### 2. slice
```go
containers := []string{"nginx", "sidecar"}
```
- Pod の `containers` や `volumes` などで頻出
- 追加や削除が多い

### 3. map
```go
labels := map[string]string{
    "app": "web",
}
```
- Kubernetesの label, annotation で非常に多い
- `map[string]string` は頻出型

### 4. pointer
```go
replicas := int32(3)
```
- Kubernetes API の多くのフィールドは pointer 型であることが多い
- `nil` と `0` の違いを理解する必要がある

### 5. for range
```go
for _, pod := range pods.Items {
    fmt.Println(pod.Name)
}
```
- リストの取り出しで超重要
- `List`, `Items` のループに必須

### 6. type alias / custom type
```go
type Namespace string
```
- Kubernetesでは型を意識して設計しやすい
- そのまま使う場面は多くないが、型安全性の観点で重要

## 4. Kubernetes API との相性が良いGoの特徴

### (1) 型安全性
- YAML の key と struct が対応しやすい
- 実行時に変な型を入れにくい

### (2) 標準ライブラリが十分
- JSON、HTTP、context、time などをそのまま使える
- Kubernetes client との相性が良い

### (3) CLI 開発に向く
- `kubectl` のようなツールは Go で実装されやすい
- `cobra` や `urfave/cli` などを使うと作りやすい

## 5. Kubernetes開発でよく見るGoの設計パターン

### controller パターン
- リソースの変化を watch する
- 変更があったら reconcile を実行する

```go
func (r *Reconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    return ctrl.Result{}, nil
}
```

### informer / cache パターン
- APIサーバーの変更をキャッシュして監視する
- controller が状態を維持する設計

### client-go パターン
- `Clientset` で Kubernetes API を叩く
- `DynamicClient` なども使える

## 6. Kubernetesで非常に大事な概念

### 1. YAML と struct の対応
- Kubernetes は YAML を中心に扱う
- Go では struct に落として扱う
- JSONタグが重要

### 2. 不変条件の考え方
- Kubernetes リソースは「現在の状態」と「望ましい状態」を比較する
- controller の設計では reconcile の考え方が基本

### 3. 冪等性
- 同じ操作を何度実行しても、最終状態が同じになる設計が重要
- Kubernetes は API の冪等性を意識する

### 4. watch と loop
- controller はイベント駆動で動く
- 無限ループや再試行制御を考える

## 7. Kubernetes 向けに Go を学ぶときのおすすめ順序

1. 基本文法
   - 変数、if、for、func
2. 型とデータ構造
   - struct、slice、map、pointer
3. errorとcontext
   - API呼び出しの基本
4. JSONとstructタグ
   - YAMLとの対応
5. packageと整理
   - 自分のコードを分ける
6. client-go の基本
   - Pod, Deployment, Namespace の取得
7. controller の考え方
   - reconcile
8. goroutine と channel
   - 並行処理の基礎
9. testing
   - API 関連のテスト

## 8. Kubernetesでよく作るGoの小さな例

### 例1: Namespace の一覧取得
```go
package main

import (
    "context"
    "fmt"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
)

func main() {
    kubeconfig := clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename()
    config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
    if err != nil {
        panic(err)
    }

    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        panic(err)
    }

    namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
    if err != nil {
        panic(err)
    }

    for _, ns := range namespaces.Items {
        fmt.Println(ns.Name)
    }
}
```

### 例2: 簡単な Kubernetes リソース取得
```go
pod, err := clientset.CoreV1().Pods("default").Get(ctx, "mypod", metav1.GetOptions{})
if err != nil {
    fmt.Println("error:", err)
    return
}
fmt.Println(pod.Name)
```

## 9. 初心者が特に注意したいこと
- `map` と `struct` の使い分けを意識する
- `nil` とゼロ値の違いを理解する
- `context` を無視しない
- `error` のハンドリングを省略しない
- `for range` の取り方を身につける
- `yaml` と `json` の変換を意識する

## 10. Kubernetes向けGo学習でのおすすめ考え方
- 「Goの文法を覚える」より「Kubernetes API を扱えるようになる」ことが大事
- まずは `client-go` を使った簡単なリソース取得から始める
- 次に controller の考え方を理解する
- その後、operator や custom resource を作る

## 11. まとめ
Kubernetes で Go を使うなら、単なるバックエンドの文法だけではなく、次の考え方が重要。

- struct と JSON タグ
- slice と map
- error と context
- client-go の使い方
- controller の reconcile 思考
- watch / cache / event-driven design

このあたりを押さえておくと、Kubernetes周りのGo開発がかなり理解しやすくなる。

## 12. 次にやるといいこと
1. `kubectl get pods` の出力を struct に落とす感覚をつかむ
2. `client-go` で Namespace / Pod を取得する
3. `yaml` と JSON を Go の struct で扱う練習をする
4. controller の基本を読む
5. その後に operator を作ってみる
