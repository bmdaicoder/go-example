package prototype_test

import (
	"go-example/patterns/design-patterns/creational/prototype"
	"testing"
)

func TestPrototype(t *testing.T) {
	person1 := &prototype.User{Name: "Minh Dai"}
	person2 := person1.Clone()

	if person1.GetInfo() != person2.GetInfo() {
		t.Errorf("These two instance do not have the same case.")
	}
}
