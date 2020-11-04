// 简单工厂模式（Simple Factory）
// 简单工厂只是为客户端生成一个实例，而不会向客户端公开任何实例化逻辑
// 在面向对象编程（OOP）中，工厂是用于创建其他对象的对象 - 正式工厂是一种函数或方法，
// 它从一些方法调用返回变化的原型或类的对象，这被假定为“新”.

package examples

import "fmt"

// desk 桌子
type desk interface {
	GetWidth() float32
	GetHeight() float32
}

// WoodenDesk 结构体
type WoodenDesk struct {
	width  float32 // 宽度
	height float32 // 长度
}

// GetWidth 获取宽
func (woodenDoor *WoodenDesk) GetWidth() float32 {
	return woodenDoor.width
}

// GetHeight 获取高度
func (woodenDoor *WoodenDesk) GetHeight() float32 {
	return woodenDoor.height
}

// MakeDesk 工厂模式创建 WoodenDesk
func MakeDesk(width float32, height float32) *WoodenDesk {
	return &WoodenDesk{width: width, height: height}
}

// SimpleFactory 简单工厂模式
func SimpleFactory() {
	desk := MakeDesk(23, 12)
	fmt.Println(desk.GetWidth())
}
