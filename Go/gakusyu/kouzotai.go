package main

import("fmt")

func main() {
	// type User struct {
	// 	Name string
	// 	Age int
	// }

	// user1 := User{Name: "山田太郎", Age: 20}
	// user2 := User{Name: "山田花子", Age: 18}
	// user3 := User{Name: "山田次郎", Age: 25}

	// fmt.Println(user1.Name, user1.Age)
	// fmt.Println(user2.Name, user2.Age)
	// fmt.Println(user3.Name, user3.Age)

	//メソッドの呼び出し
p := Person{Name: "山田太郎", Age: 20}
value := p.SelfIntroduction()
fmt.Println(value)

}

//振る舞いの定義(メソッド)

//type Person struct {
		Name string
		Age int
}
//func (p Person) SelfIntroduction() string {
		return "私の名前は" + p.Name + fmt.Sprint(p.Age) + "歳です。"
}