// +build ignore

package main

import "fmt"

type Vertex struct {
	X, Y int
}

type Vertex2 struct {
	X, Y int
}

var (
	a = Vertex2{1, 2}  // 类型为 Vertex2
	b = Vertex{X: 1}  // Y:0 被省略
	c = Vertex{}      // X:0 和 Y:0
	d = &Vertex{1, 2} // 类型为 *Vertex
)

func main() {
	fmt.Printf("%T(%v) \n", a, a);
	fmt.Printf("%T(%v) \n", b, b);
	fmt.Printf("%T(%v) \n", c, c);
	fmt.Printf("%T(%v) \n", d, d);
}
