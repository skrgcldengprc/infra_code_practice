# Goをマスターしたらできること（インフラ・クラウドネイティブ編）

## 1. この資料の前提
- 主な目標は「Kubernetes・インフラ・クラウドネイティブの自動化とツール開発」
- Go は Terraform, Kubernetes, Argo, Prometheus 関連のツール群で広く使われている
- 学習の価値は「手動作業の自動化」「API 連携」「CLI開発」「コントローラー開発」へ広がる

## 2. Goを覚えるとインフラ周りで何ができるか

### 2-1. Kubernetesの自動化
- Pod, Deployment, Service, Namespace をプログラムから作成・更新・削除
- 監視して変化を検知し、自動修正する controller を作る
- Kubernetes の設定をコード化して管理する

必要なGoの知識:
- struct
- JSON/YAML
- map / slice
- client-go
- context
- error handling

例:
```go
pod, err := clientset.CoreV1().Pods("default").Get(ctx, "app-pod", metav1.GetOptions{})
if err != nil {
    return err
}
fmt.Println(pod.Name)
```

---

### 2-2. Kubernetes Operator を作る
- 独自リソースを管理できる
- 自動修復・自動スケーリング・自動デプロイが作れる
- アプリとKubernetesをつなぐ制御ロジックを実装できる

必要なGoの知識:
- controller-runtime
- reconcile
- watch
- informer
- custom resource definition（CRD）
- ループと再試行の設計

例:
```go
func (r *MyReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    // 望ましい状態と現在状態を比較する
    return ctrl.Result{}, nil
}
```

---

### 2-3. Terraform/Infrastructure as Code の補助ツールを作る
- テンプレート生成
- バリデーション処理
- 実行ログの整形
- Terraform の設定生成支援CLI

必要なGoの知識:
- file I/O
- JSON / YAML
- string 操作
- CLI実装
- 標準ライブラリの os, filepath

例:
```go
content, err := os.ReadFile("main.tf")
if err != nil {
    panic(err)
}
fmt.Println(string(content))
```

---

### 2-4. Cloud CLI や自動運用ツールを作る
- AWS, GCP, Azure の API を叩く小さなツール
- 監視メトリクスの収集
- リソース状態の確認ツール
- デプロイ前の要件チェック

必要なGoの知識:
- HTTP client
- JSON decode
- struct tag
- context
- retry / timeout

例:
```go
resp, err := http.Get("https://api.example.com/health")
if err != nil {
    return err
}
defer resp.Body.Close()
```

---

### 2-5. CI/CD パイプライン用ツールを作る
- GitHub Actions などの前段後段で使う補助ツール
- デプロイ対象の判定
- 変更ファイルの収集
- デプロイ前の検証

必要なGoの知識:
- os.Args
- file walk
- regexp
- time
- JSON
- CLIの引数処理

例:
```go
func main() {
    fmt.Println("引数:", os.Args[1:])
}
```

---

### 2-6. Prometheus / メトリクス周りのツール開発
- メトリクス収集スクリプト
- 監視対象の状態確認
- alert rule 生成支援
- メトリクス変換処理

必要なGoの知識:
- HTTP server / client
- time
- fmt
- struct
- goroutine

例:
```go
go func() {
    for range time.Tick(5 * time.Second) {
        fmt.Println("metrics check")
    }
}()
```

---

### 2-7. ログ収集・監視・トラブルシュート支援ツール
- ログの抽出と集計
- 監視対象の状態サマリ
- エラーの自動分類
- 通知用のテキスト生成

必要なGoの知識:
- file I/O
- strings
- regexp
- map
- slice
- JSON

例:
```go
if strings.Contains(line, "ERROR") {
    fmt.Println("エラー行:", line)
}
```

---

### 2-8. GitOps や自動構成ツールの開発
- Git リポジトリの状態を見て構成を作る
- 差分を検出して適用対象を決める
- 自動生成された manifest の整形

必要なGoの知識:
- file I/O
- YAML parsing
- exec コマンド実行
- os/exec
- CLI設計

例:
```go
cmd := exec.Command("kubectl", "get", "pods", "-A")
output, err := cmd.Output()
if err != nil {
    panic(err)
}
fmt.Println(string(output))
```

---

## 3. 技術ごとに必要になるGoの構文・考え方

### Terraform系
必要な構文・概念:
- string / map / slice
- file read/write
- JSON/YAML
- os, filepath
- CLI引数

使う場面:
- tfvars の生成
- HCL テンプレートの補助
- 事前チェックツール

---

### Kubernetes系
必要な構文・概念:
- struct
- pointer
- for range
- map
- context
- error
- client-go
- controller-runtime

使う場面:
- Pod 管理
- Deployment 作成
- custom controller
- operator

---

### Cloud API系
必要な構文・概念:
- HTTP client
- JSON decode
- time
- context
- retry
- struct tag

使う場面:
- AWS/GCP/Azure API呼び出し
- リソース管理ツール
- 監視や運用支援

---

### CI/CD系
必要な構文・概念:
- os.Args
- file I/O
- regexp
- exec.Command
- time
- error handling

使う場面:
- ビルド前チェック
- デプロイ対象判定
- 自動化スクリプト

---

### Observability系
必要な構文・概念:
- goroutine
- time.Tick
- HTTP server
- map
- strings / regexp

使う場面:
- ログ監視
- メトリクス収集
- サービスランナー

---

## 4. Goをマスターしたらできる代表的な実務例

### 4-1. Kubernetes Custom Controller
- 外部システムの状態を監視して、Kubernetes リソースを更新する
- 独自の CRD と紐付いた制御ロジックを実装する

### 4-2. Kubernetes CLI ツール
- namespace 一覧を見やすく整形
- pod の状態を集計して出力
- 変更対象を自動検出してデプロイ対象にする

### 4-3. インフラ自動化ツール
- Terraform の plan/apply をsafe化する補助ツール
- 環境差分の検出とレポート出力
- 削除対象の確認と制御

### 4-4. Cloud 自動運用ツール
- リソースの状態確認
- 作成と削除の自動化
- 事前チェックと安全性制御

### 4-5. GitOps 補助ツール
- YAML 生成
- 差分取得
- デプロイ対象の自動選定

---

## 5. まず覚えるべきGoの要点（インフラ向け）
- `struct` と `json` タグ
- `slice` と `map`
- `for range`
- `error` と `context`
- pointer と nil
- file I/O
- http client
- goroutine
- CLI 引数

この辺を覚えると、インフラ周りの自動化ツールを作り始めやすい。

---

## 6. まとめ
Goは「インフラ・クラウドネイティブ分野」で特に強い言語である。

- Kubernetes を支える実装の中心言語
- 自動化ツール開発に最適
- CLI・controller・operator・API連携が扱いやすい
- インフラ運用の効率化に直結する

特にKubernetesを作りたいなら、Go の次の分野を優先して学ぶと実務に直結する。

1. struct + JSON/YAML
2. client-go
3. context + error
4. controller-runtime
5. goroutine と並行処理
6. CLI開発

## 7. 次にやるべきこと
- Kubernetes API の簡単な取得コードを書く
- Pod や Namespace を Go で一覧取得してみる
- `client-go` の基本例を最小構成で動かす
- `controller-runtime` の reconcile の流れを読む
- その後、自分の小さな controller や CLI を作る
