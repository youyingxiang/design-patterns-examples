package main

import (
	"./examples"
)

func main() {
	// 工厂模式例子
	examples.SimpleFactory()

	// 工厂方法模式（Factory Method）
	examples.FactoryMethod()

	// 抽象工厂模式（Abstract Factory）
	examples.AbstractFactory()

}
