package main

import (
	"fmt"
	"math/rand"
)

func main() {
	count := 0 //短声明

	for count < 10 {
		var num = rand.Intn(10) + 1
		fmt.Println(num)

		count++
	}
	fmt.Println("---------------------")
	shortVar()
}

func shortVar() {

	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}
}

// 在 if 里使用短生命来声明变量
func fooFor() {
	if num := rand.Intn(3); num == 0 {
		fmt.Println("Space Adventures")
	} else if num == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
	}
}

func fooSwitch() {
	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("Space Adventure")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	case 3:
		fmt.Println("Random spaceline #", num)
	}
}
