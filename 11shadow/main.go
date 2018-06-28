package main

import "fmt"

func main() {
	i := 3
	{
		i := 42
		fmt.Println(i)
	}
	fmt.Println(i)
}
