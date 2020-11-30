package main

import "fmt"

func swap(x string, y string) (string, string) {
	fmt.Println(x)
	return x, y
}

func main() {
	a, b := swap("Hello", "World")
	fmt.Println(a, b)
}
