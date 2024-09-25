package singleton_test

import (
	"fmt"
	"go-example/patterns/creational/singleton"
	"testing"
)

func TestSingleton(t *testing.T) {
	call1 := singleton.GetInstance()
	call2 := singleton.GetInstance()

	fmt.Println(&call1, &call2)

	if call1 != call2 {
		t.Errorf("These two instance do not have the same case.")
	}
}
