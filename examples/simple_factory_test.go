package examples

import (
	"testing"
)

// command: go test -v simple_factory_test.go simple_factory.go
func TestSimpleFactory(t *testing.T) {
	desk := MakeDesk(24, 12)

	if desk.GetWidth() != 24 || desk.GetHeight() != 12 {
		t.Fail()
	}

}
