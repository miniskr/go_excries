package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	//var s, sep string
	//s, sep := "", ""

	// 1.改良写法
	// for _, arg := range os.Args[1:] {
	// 	s += sep + arg
	// 	sep = " "
	// }

	// 0.基本写法
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }

	// fmt.Println(s)

	// 2. join写法
	fmt.Println(strings.Join(os.Args[1:], " "))
}
