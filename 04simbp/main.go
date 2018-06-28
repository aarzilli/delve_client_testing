package main

import (
	"sync"
	"fmt"
)

func coroutine(i int, wg *sync.WaitGroup) {
	fmt.Println("hello from coroutine ", i)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go coroutine(i, &wg)
	}
	wg.Wait()
}
