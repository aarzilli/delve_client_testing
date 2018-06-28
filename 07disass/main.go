package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.Breakpoint()
	fmt.Println("hello")
	fmt.Println("goodbye")
}