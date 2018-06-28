package main

import (
	"fmt"
	"sync"
	"runtime"
)

func coroutine(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	
	j := i * 2
	
	if i == 99 {
		runtime.Breakpoint()
	}
	
	fmt.Println("hello ", i, j)
	fmt.Println("goodbye", i, j)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go coroutine(i, &wg)
	}
	wg.Wait()
}
