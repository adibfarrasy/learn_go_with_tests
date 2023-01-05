package main

import "fmt"

func main() {
	num := []int{1, 1, 2, 3, 5, 8, 13}
	for _, item := range num {
		fmt.Println(item)
	}
}
