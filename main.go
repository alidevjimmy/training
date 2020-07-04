package main

import (
	"fmt"
	"math/cmplx"
)

var (
	// ToBe is just for test
	ToBe bool = true
	// MaxInt is Just for test
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	i      int8       = 50
)

// Pi is pi number
const Pi float32 = 3.14

func main() {
	i := 20
	if i == 20 {
		fmt.Println("ok")
	}

	// fmt.Printf("%T\n", i)
}
