package hello

import "testing"

func TestHello(t *testing.T) {
	result := Hello()
	if result != "Hello, world." {
		t.Fail()
	}
}
