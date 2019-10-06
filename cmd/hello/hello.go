package main

import (
	"fmt"
	"time"

	customHello "github.com/stoksc/service-watcher/internal/hello"
	"github.com/stoksc/service-watcher/pkg/hello"
)

func main() {
	for {
		fmt.Println(customHello.CustomHello("stoksc"))
		fmt.Println(hello.Hello())
		time.Sleep(time.Second)
	}
}
