package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
}

func (p Person) SelfIntroduction() string{
	return "私の名前は" + p.Name + "です。"
}

func main(){
	p := Person{Name: "テスト二郎", Age: 30}
	value := p.SelfIntroduction()
	fmt.Println(value)
}