// 简单工厂模式（Simple Factory）
// 简单工厂只是为客户端生成一个实例，而不会向客户端公开任何实例化逻辑
// 在面向对象编程（OOP）中，工厂是用于创建其他对象的对象 - 正式工厂是一种函数或方法，
// 它从一些方法调用返回变化的原型或类的对象，这被假定为“新”.

package examples

// 首先，我们有一个门界面和实现 .

type door interface {
	getWidth() float32
	getHeight() float32
}

// WoodenDoor 结构体
type WoodenDoor struct {
	width  float32 // 宽度
	height float32 // 长度
}

// GetWidth 获取宽
func (woodenDoor *WoodenDoor) GetWidth() float32 {
	return woodenDoor.width
}

// GetHeight 获取高度
func (woodenDoor *WoodenDoor) GetHeight() float32 {
	return woodenDoor.height
}

// MakeDoor 工厂模式创建 WoodenDoor
func MakeDoor(width float32, height float32) *WoodenDoor {
	return &WoodenDoor{width: width, height: height}
}
