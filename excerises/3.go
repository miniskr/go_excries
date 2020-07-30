/*
* 实现一个猜数游戏，首先定义一个1-100的整数，
* 然后让计算机生成一个1-100的随机数，并显示
* 计算机猜测的结果是太大了还是太小了，没猜对
* 就继续猜，直到猜对为止
 */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	x := 0
	count := 0
	fmt.Println("请输入一个整数:")

	fmt.Scanln(&x)

	fmt.Println("You have printed: %v", x)

	for {
		guess := rand.Intn(101)
		fmt.Println("Computer guess %v", guess)
		if guess == x {
			break
		}
		count++
		time.Sleep(time.Millisecond * 100)
	}
	fmt.Println("电脑猜了: %v 次", count)
}
