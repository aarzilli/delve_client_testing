package main

import (
	"sync"
	"fmt"
	"time"
)

func coroutine(i int, wg *sync.WaitGroup) {
	for j := 0; j < 10; j++ {
		time.Sleep(10 * time.Millisecond)
		fmt.Println("hello from coroutine ", i, j)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go coroutine(i, &wg)
	}
	fmt.Println("setup done")
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("loop ", i)
	}
	wg.Wait()
}
