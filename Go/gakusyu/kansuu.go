package main

import "fmt"

//mainという関数は特殊な意味を持ち、プログラムが実行された際、最初に実行される関数となる。

func main() {
	display()
	result := add(100, 200, 200)
	fmt.Println(result)
}

func display() {
	fmt.Println("HelloWorld!!")
}

func add(value1 int, value2 int, value3 int) int {
	//result := value1 * value2 / value3
	//fmt.Println(result)
	return value1 + value2 + value3
}
