package main

import (
	"fmt"
	"reflect"
	"structs_methods_interfaces/service"
)

func main() {
	fmt.Println(check(new(service.Rectangle), new(service.Shape)))
}

func check(n interface{}, i interface{}) bool {
	ti := reflect.TypeOf(i).Elem()
	return reflect.TypeOf(n).Implements(ti)
}
