package main

import (
	"os"

	"dependency_injection/di"
)

func main() {
	di.Greet(os.Stdout, "Adib")
}
