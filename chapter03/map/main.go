package main

import (
	"fmt"
)

func main() {
	scene := make(map[string]int)

	scene["route"] = 66
	scene["brail"] = 4
	scene["china"] = 86

	// 无序
	for k, v := range scene {
		fmt.Println(k, v)
	}


	// 排序
	// 声明一个切片保存map数据
	var scene List []string

	// 将map数据遍历复制到切片中
	for k := range scene{
		scene List = append(scene List, k)
	}

	// 对切片进行排序
	sort.Strings(scene List)

	// 输出
	fmt.Println(scene List)
}
