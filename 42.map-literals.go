// +build ignore

// map 的文法
package main

import "fmt"

type V3 struct {
	x, y float64
}

var i = map[string] V3 {
	// 如果结构体内的值都在一行，最后一个值后面可以忽略,
	"a": V3 { 40.68433, -74.39967 },
	// 如果结构体内的值不在同一行，最后一个值后面必须,
	"b": V3 {
		37.42202, -122.08408,
	},
}

// 如果顶级类型只有类型名的化，可以在文法元素中省略键名
var j = map[string] V3 {
	"a": { 40.68433, -74.39967 },
	"b": { 37.42202, -122.08408 },
}

func main() {
	fmt.Println(i)
	fmt.Println(j)
}
