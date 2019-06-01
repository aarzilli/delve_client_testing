package main

import (
	"unsafe"
	"reflect"
	"fmt"
	"sync"
	"runtime"
)

func done() {
	buf := make([]byte, 100000000)
	slh := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	var s string
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	sh.Data = slh.Data
	sh.Len = slh.Len
	fmt.Printf("about to stop\n")
	runtime.Breakpoint()
	fmt.Printf("this\n")
	fmt.Printf("should\n")
	fmt.Printf("be\n")
	fmt.Printf("fast\n")
}

func makeStack(depth int) {
	if depth == 0 {
		done()
		return
	}
	makeStack(depth-1)
}

func makeGoroutines(i int, c *sync.Cond) {
	c.L.Lock()
	c.Wait()
	c.L.Unlock()
}

func main() {
	var mu sync.Mutex
	c := sync.NewCond(&mu)
	for i := 0; i < 1000000; i++ {
		go makeGoroutines(i, c)
	}
	makeStack(10000000)
	c.Broadcast()
}
