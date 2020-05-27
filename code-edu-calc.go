package main

import (
	"fmt"
)

func sum(x int, y int) (result int) {
	result = x + y
	return result
}

func main() {
	fmt.Println("Code.education Sum!")
	result := sum(5, 5);
	fmt.Printf("5 + 5 = %d", result);
}