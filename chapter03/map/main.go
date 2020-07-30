package main

import "fmt"

func main() {
	scene := make(map[string]int)

	scene["route"] = 66

	fmt.Println(scene["route"])

	v := scene["route1"]

	fmt.Println(v)
}
