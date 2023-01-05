package main

import (
	"fmt"
	hello "hello/service"
)

func main() {
	name := "adib"
	fmt.Println(hello.Hello(&name, nil))
}
