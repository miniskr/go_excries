package main

import (
	"fmt"
)

// 指针
func main() {

	// Go语言使用&操作符放在变量前进行取地址操作
	// ptr := &v //v的类型为T

	var cat int = 1
	var str string = "banana"

	fmt.Printf("%p %p", &cat, &str)
}
