package examples

import "testing"

// command: go test -v abstract_factory_test.go abstract_factory.go
func TestAbstractFactory(t *testing.T) {
	// 木门工厂
	woodenDoorFactory := &WoodenDoorFactory{}
	// 木门的描述
	if str := woodenDoorFactory.MakeDoor().GetDescription(); str != "我是木门！" {
		t.Fail()
	}
	// 木匠的描述
	if str := woodenDoorFactory.MakeFittingExpert().GetDescription(); str != "我是装木门的人！" {
		t.Fail()
	}

	// 铁门工厂
	ironDoorFactory := &IronDoorFactory{}
	// 铁门的描述
	if str := ironDoorFactory.MakeDoor().GetDescription(); str != "我是铁门！" {
		t.Fail()
	}
	// 铁匠的描述
	if str := ironDoorFactory.MakeFittingExpert().GetDescription(); str != "我是装铁门的人！" {
		t.Fail()
	}
}
