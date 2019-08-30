// +build ignore

package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) add() float64 {
	return v.X + v.Y
}

func (v *Vertex) decrease() float64 {
	return v.X - v.Y
}

func main() {
	res := &Vertex{3, 4}
	fmt.Println(res)
	fmt.Println(res.add())
	fmt.Println(res.decrease())
}
