// 抽象工厂模式（Abstract Factory）
// 现实世界的例子:从 Simple Factory 扩展我们的门例子。根据您的需求，您可以从木门店，铁门的铁门或相关商店的 PVC 门获得木
// 门。另外，你可能需要一个有不同种类特色的家伙来安装门，例如木门木匠，铁门焊机等。你可以看到门之间存在依赖关系，木门需要木匠，铁门需要焊工等
// 简单来说; 将个人但相关/依赖工厂分组在一起而不指定其具体类别的工厂。
// 维基百科说:抽象工厂模式提供了一种封装一组具有共同主题但没有指定其具体类的单个工厂的方法

package examples

import "fmt"

// door 接口
type door interface {
	// GetDescription 获取描述信息
	GetDescription()
}

// WoodenDoor 木门
type WoodenDoor struct {
}

// GetDescription 木门的描述
func (woodenDoor *WoodenDoor) GetDescription() {
	fmt.Println("我是木门！")
}

// IronDoor 铁门
type IronDoor struct {
}

// GetDescription 铁门的描述
func (ironDoor *IronDoor) GetDescription() {
	fmt.Println("我是铁门！")
}

// doorFittingExpert 接口（装门的人）
type doorFittingExpert interface {
	// GetDescription 获取描述信息
	GetDescription()
}

// Welder 装铁门的人
type Welder struct {
}

// GetDescription 装铁门的人的描述
func (welder *Welder) GetDescription() {
	fmt.Println("我是装铁门的人！")
}

// Carpenter 装木门的人
type Carpenter struct {
}

// GetDescription 装木门的人的描述
func (carpenter *Carpenter) GetDescription() {
	fmt.Println("我是装木门的人！")
}

// doorFactory 现在我们有抽象工厂，让我们制作相关对象的工厂，即木门工厂将创建一个木门和木匠，铁门工厂将创建一个铁门和铁匠
type doorFactory interface {
	// MakeDoor 创建门
	MakeDoor()
	// MakeFittingExpert 选择工匠
	MakeFittingExpert()
}

// WoodenDoorFactory 木门工厂
type WoodenDoorFactory struct {
}

// MakeDoor 创建木门
func (woodenDoorFactory *WoodenDoorFactory) MakeDoor() *WoodenDoor {
	return &WoodenDoor{}
}

// MakeFittingExpert 选择木匠
func (woodenDoorFactory *WoodenDoorFactory) MakeFittingExpert() *Carpenter {
	return &Carpenter{}
}

// IronDoorFactory 铁门工厂
type IronDoorFactory struct {
}

// MakeDoor 创建铁门
func (ironDoorFactory *IronDoorFactory) MakeDoor() *IronDoor {
	return &IronDoor{}
}

// MakeFittingExpert 选择木匠
func (ironDoorFactory *IronDoorFactory) MakeFittingExpert() *Welder {
	return &Welder{}
}

// AbstractFactory 抽象工厂模式
func AbstractFactory() {
	// 木门工厂
	woodenDoorFactory := &WoodenDoorFactory{}
	// 木门的描述
	woodenDoorFactory.MakeDoor().GetDescription()
	// 木匠的描述
	woodenDoorFactory.MakeFittingExpert().GetDescription()

	// 铁门工厂
	ironDoorFactory := &IronDoorFactory{}
	// 铁门的描述
	ironDoorFactory.MakeDoor().GetDescription()
	// 铁匠的描述
	ironDoorFactory.MakeFittingExpert().GetDescription()

}
