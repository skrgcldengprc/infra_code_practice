package main

import "fmt"

func main() {
	fmt.Println("問題4: 配列とスライス")
	fmt.Println("1. 3つのサーバー名を持つスライスを作成する")
	fmt.Println("2. そのスライスを表示する")
	fmt.Println("3. 先頭に 'web-01' を追加して、表示する")

	servers := []string{"web-02/", "web-03/", "web-04/", "web-05/"}
	servers = append([]string{"web-01/"}, servers...)
	fmt.Println(servers)
}

// TODO: サーバー名のスライスを作成し、追加してください。
// servers := []string{"web-02", "app-01", "db-01"}
// servers = append([]string{"web-01"}, servers...)
// fmt.Println(servers)
