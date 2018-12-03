package main

import (
	"fmt"

	hello "github.com/notsu/gomono/services/post/hello"
	helloV4 "github.com/notsu/gomono/services/post/v4/hello"
	helloV5 "github.com/notsu/gomono/services/post/v5/hello"
	helloV501 "github.com/notsu/gomono/services/post/v501/hello"
)

func main() {
	fmt.Println("Hello from USER")
	fmt.Println(hello.Hello())
	fmt.Println(helloV4.Hello())
	fmt.Println(helloV5.Hello())
	fmt.Println(helloV501.Hello())
}
