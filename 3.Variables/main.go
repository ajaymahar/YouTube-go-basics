package main

import (
	"fmt"
)

// var (
// 	name   string = "Aj"
// 	age    int    = 29
// 	amount int    = 99
// )

var (
	i = j
	j = 20
)

func main() {
	// 	var i = 10
	// 	fmt.Println(i)
	//
	// 	var j = 99.9
	// 	fmt.Println(j)

	// var amount, discount int = 100, 90
	// fmt.Println(amount, discount)

	// var server, ip, isAlive = "localhost", "127.0.0.1", true
	// fmt.Printf("server=[%v], ip=[%v], isAlive=[%v]\n", server, ip, isAlive)

	// age := 28
	// fmt.Println(age)

	// name, email := "AJ", "aj@email.com"
	// fmt.Println(name, email)

	// name, age := "AJ", 28
	// fmt.Println(name, age)

	//Scope package level var declaration
	// fmt.Println(name)

	//block level var declaration
	// var name string = "AM"
	// fmt.Println(name)
	// msg()

	// 	var i int
	// 	var j float32
	// 	var b bool
	// 	var str string
	// 	var arr [5]int
	// 	var sl []int
	// 	var p *int
	// 	var m map[int]string
	//
	// 	fmt.Println(i)
	// 	fmt.Println(j)
	// 	fmt.Println(b)
	// 	fmt.Println(str)
	// 	fmt.Println(arr)
	// 	fmt.Println(sl)
	// 	fmt.Println(p)
	// 	fmt.Println(m)
	//

	// 	{
	// 		var name = "Mr. X"
	// 		fmt.Println(name)
	//
	// 	}
	// 	fmt.Println(name)

	// i := 7
	// j := 19
	// fmt.Println(j)

	// j := 19
	// _ = 7
	// fmt.Println(j)

	// amount, _ := someFunc()
	// fmt.Println(amount)

	// fmt.Println(name)
	// fmt.Println(age)
	// fmt.Println(amount)

	// i := 10
	// fmt.Println(i)

	i, j := j, 10
	fmt.Println(i + j)
}

// func someFunc() (int, error) {
// 	return 10, errors.New("Error")
// }

// func msg() {
// 	fmt.Println(name)
// }
