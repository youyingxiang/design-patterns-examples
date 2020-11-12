// 原型模式（Prototype）
// 原型模式是软件开发中的创新设计模式。当要创建的对象类型由原型实例确定时使用它，该实例被克隆以生成新对象。
// 模式特点：用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象。

package examples

import "fmt"

// Sheep 🐑
type Sheep struct {
	name     string
	category string
}

// SetName ...
func (sp *Sheep) SetName(name string) {
	sp.name = name
}

// GetName ...
func (sp *Sheep) GetName() string {
	return sp.name
}

// SeCategory ...
func (sp *Sheep) SeCategory(category string) {
	sp.category = category
}

// GetCategory ...
func (sp *Sheep) GetCategory() string {
	return sp.category
}

// Clone 克隆一个一摸一样的 🐑
func (sp *Sheep) Clone() *Sheep {
	clonesp := *sp
	return &clonesp
}

// GetDescription ...
func (sp *Sheep) GetDescription() {
	fmt.Printf(" 🐑 的名字：%v ;🐑 的种类：%v \n", sp.name, sp.category)
}

// NewSheep 工厂方法创建一只 🐑
func NewSheep() *Sheep {
	return &Sheep{}
}
