package hello

import "testing"

func TestCustomHello(t *testing.T) {
	name := "stoksc"
	result := CustomHello(name)
	if result != "Hello, stoksc" {
		t.Fail()
	}
}
