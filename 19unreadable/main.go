package main

import (
	"reflect"
	"unsafe"
	"fmt"
	"runtime"
)

func main() {
	var s string
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sh.Data = 0x1
	sh.Len = 100
	runtime.Breakpoint()
	fmt.Printf("%s\n", s)
}
