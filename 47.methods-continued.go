// +build ignore

package main

import (
	"fmt"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	// 取正，示例
	a := MyFloat(11.000)
	b := MyFloat(-22.000)
	fmt.Println(a.Abs())
	fmt.Println(b.Abs())
}
