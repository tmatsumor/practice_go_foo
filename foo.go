package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func main() {
	result := Add(1, 2)
	fmt.Printf(`Add(1, 2) is `+"%v", result)
}
