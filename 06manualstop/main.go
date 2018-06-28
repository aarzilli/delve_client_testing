package main

import (
	"time"
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("I feel asleep")
	}
}
