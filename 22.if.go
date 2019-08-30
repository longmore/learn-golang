// +build ignore

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	const f = "%T(%v)\n"
	fmt.Printf(f, x, x)
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
