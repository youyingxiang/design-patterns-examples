package examples

import "testing"

// command: go test -v prototype_test.go prototype.go
func TestPrototype(t *testing.T) {
	t.Log("创建一只 🐑")
	sheep := NewSheep()
	sheep.SetName("Jolly")
	sheep.SeCategory("山地绵羊")

	t.Log("克隆一只 🐑")
	clonesheep := sheep.Clone()
	clonesheep.SetName("Dolly")
	clonesheep.SeCategory("澳大利亚绵羊")

	if sheep.GetName() != "Jolly" || sheep.GetCategory() != "山地绵羊" {
		t.Fail()
	}

	if clonesheep.GetName() != "Dolly" || clonesheep.GetCategory() != "澳大利亚绵羊" {
		t.Fail()
	}

}
