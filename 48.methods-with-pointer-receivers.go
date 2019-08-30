// +build ignore

// 接收者为指针的方法
/**
使用场景：
1,避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）。
2,方法可以修改接收者指向的值。
*/

package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Add() float64 {
	return v.X + v.Y
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v, v.Add())
}
