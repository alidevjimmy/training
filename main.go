package main

import (
	"fmt"
	"math"
)

func swap(first, second string) (string, string) {
	return second, first
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b, ' ', math.Pi)
}
