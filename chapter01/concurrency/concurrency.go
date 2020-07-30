package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数据生产者
func producer(header string, channel chan<- string) {

	// 无限循环产生数据
	for {
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())

		time.Sleep(time.Second * 2)
	}
}

// 数据消费者
func consumer(channel <-chan string) {
	// 不停的获取数据
	for {
		message := <-channel

		fmt.Println(message)
	}
}

func main() {
	// 创建一个字符串类型通道
	channel := make(chan string)

	// 创建producer函数的并发
	go producer("cat", channel)
	go producer("dog", channel)

	// 数据消费者
	consumer(channel)
}
