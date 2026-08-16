package main

import "fmt"

// func main() {
// 	practice := "練習"
// 	fmt.Println("私の名前は" + practice + "です。")

// 	var base float32 = 10
// 	var height float32 = 20
// 	var two float32 = 2
// 	fmt.Println(base * height / two)
// }

// func main() {
// 	value := 11

// 	if value %2 == 0 {
// 	    fmt.Println("valueは偶数です。")
// 	}else if value %2 != 0 {
// 	    fmt.Println("valueは奇数です。")
// 	}
// }

// func main() {
// 	year := 2024

// 	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
// 		fmt.Println("閏年です。")
// 	} else {
// 		fmt.Println("通常の年です")
// 	}
// }

// func main() {
// 	week := 6

// 	switch week {

// 	case 1:
// 		fmt.Println("月曜日です。")
// 	case 2:
// 		fmt.Println("火曜日です。")
// 	case 3:
// 		fmt.Println("水曜日です。")
// 	case 4:
// 		fmt.Println("木曜日です。")
// 	case 5:
// 		fmt.Println("金曜日です。")
// 	case 6:
// 		fmt.Println("土曜日です。")
// 	case 7:
// 		fmt.Println("日曜日です。")
// 	default:
// 		fmt.Println("範囲外です。")
// 	}
// }

// func main() {

// 	for i := 0; i <= 20; i++ {
// 		fmt.Println(i, "です。")
// 	}
// }

// func main() {
// 	for i := 1; i <= 100; i++ {
// 		if i%3 == 0 && i%5 == 0 {
// 			fmt.Println("FizzBuzz")
// 		} else if i%3 == 0 {
// 			fmt.Println("Fizz")
// 		} else if i%5 == 0 {
// 			fmt.Println("Buzz")
// 		} else {
// 			fmt.Println(i)
// 		}
// 	}
// }

// func main() {
// 	result := getBMI(62,1.7)
// 	fmt.Println(result)
// }

// func getBMI(weight float32, height float32) float32{
//     return weight / (height * height)
// }

// func main(){
// 	var names [5] string
// 	names[0] = "gi"
// 	names[1] = "t"
// 	names[2] = "h"
// 	names[3] = "u"
// 	names[4] = "b"

// 	// names := [5]string{"さ","く","ら","う","ゆ"}

// 	for _, name := range names {
// 		fmt.Println(name)
// 	}
// }

func main() {
	names := []string{"r", "e", "n", "s", "y", "u"}
	names = append(names, "!")

	for _, name := range names {
		fmt.Println(name)
	}
}
