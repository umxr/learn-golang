package main

import "fmt"

var i, j int = 1, 2

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(i, j))
}
