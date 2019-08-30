// +build ingore

package main

import "fmt"

type V1 struct {
	x, y float64
}

type V2 struct {
	x, y int
}

// var m map[string]Vertex

func main() {

	var m = make(map[string]V2) // string 为 key的类型 , V2为value的结构体
	fmt.Println(m)
	m["a"] = V2{
		11,
		22,
	}
	m["b"] = V2{
		33,
		44,
	}
	fmt.Println(m)
	fmt.Println(m["a"])
	fmt.Println(m["a"].x)
	fmt.Println(m["a"].y)

}