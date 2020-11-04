package main

import (
	"fmt"

	"./examples"
)

func main() {
	// 工厂模式例子
	door := examples.MakeDoor(23, 12)
	fmt.Println(door.GetWidth())
}
