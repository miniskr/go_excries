package main

import "fmt"

func getData() (int, int) {
	return 100, 200
}

// 匿名变量
func main() {
	a, _ := getData()

	_, b := getData()

	fmt.Println(a, b)
}
