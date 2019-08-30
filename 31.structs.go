// +build ignore

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	a := Vertex{1, 2}
	fmt.Println(a)
	fmt.Println(a.X)
	fmt.Println(a.Y)
	a.X = 100 // 修改值
	fmt.Println(a)

	b := &a // 生成指针
	fmt.Println(b) // 指针地址
	fmt.Println(*b) // 指针指向值
	b.X = 200
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(a)

}

