package main

import (
	"fmt"

	"./examples"
)

func main() {
	// 工厂模式例子
	door := examples.MakeDoor(23, 12)
	fmt.Println(door.GetWidth())

	// 工厂方法模式（Factory Method）
	developer := examples.Developer{}

	developmentManager := examples.NewHiringManager{Interviewer: developer}
	developmentManager.TakeInterview()
	// developmentManager.TakeInterview()
}
