package main

import (
	"fmt"
	"runtime"
)

var v = "irrelevant"

func main() {
	runtime.Breakpoint()
	fmt.Println("hello")
	fmt.Println("goodbye", v)
}
