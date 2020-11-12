package examples

import (
	"testing"
)

// TestBuilder 测试
// command: go test -v builder_test.go builder.go
func TestBuilder(t *testing.T) {

	t.Log("制作第一个 🍔")
	burger1 := NewBurgerBuilder().AddPepperoni("火腿1").AddCheese("奶酪1").AddLettuce("生菜1").AddTomato("番茄1").Build()
	burger1.GetDescription()

	t.Log("制作第二个 🍔")
	burger2 := NewBurgerBuilder().AddPepperoni("火腿2").AddCheese("奶酪2").AddLettuce("生菜2").AddTomato("番茄2").Build()
	burger2.GetDescription()

	t.Log("制作第三个 🍔")
	burger3 := NewBurgerBuilder().AddPepperoni("火腿3").AddLettuce("生菜3").AddTomato("番茄3").Build()
	if burger3 != nil {
		t.Fail()
	}
}
