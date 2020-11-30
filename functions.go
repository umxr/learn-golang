package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Result", add(1, 2))
}
