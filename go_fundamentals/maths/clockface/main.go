package main

import (
	"os"
	"time"

	. "maths"
)

func main() {
	t := time.Now()
	SVGWriter(os.Stdout, t)
}
