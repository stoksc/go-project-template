package hello

import "fmt"

// CustomHello says hello, in a custom way.
func CustomHello(name string) string {
	return fmt.Sprintf("Hello, %v", name)
}
