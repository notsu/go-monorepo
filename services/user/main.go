package main

import (
	"fmt"

	hello "github.com/notsu/gomono/services/post/hello"
)

func main() {
	fmt.Println("Hello from USER")
	fmt.Println(hello.Hello())
}
