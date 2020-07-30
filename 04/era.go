package main

import (
	"fmt"
	"math/rand"
)

// 这个变量拥有package级别的作用域，这里不可以使用短声明声明变量
var era = "AD"

func main() {
	year := 2018

	switch month := rand.Intn(12) + 1; month {
	case 2:
		day := rand.Intn(28) + 1
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)
	}
}
