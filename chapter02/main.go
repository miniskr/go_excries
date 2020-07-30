package main

import "fmt"

func varableSwitch() {
	// 多个变量同时赋值
	// 1. 传统做法
	// var a int = 100
	// var b int = 200
	// var temp int

	// temp = a
	// a = b
	// b = temp

	// 2. 高级做法 通过异或操作
	// var a int = 100
	// var b int = 200

	// a = a ^ b
	// b = b ^ a
	// a = a ^ b

	// fmt.Println(a, b)

	//3. 简单做法
	var a int = 100
	var b int = 200

	a, b = b, a
	fmt.Println(a, b)
}

func main() {
	varableSwitch()
}
