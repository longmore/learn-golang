// +build ignore

package main

import "fmt"

func main() {
	// demo1
	a := 42
	b := &a // 生成指针，并指向a的内存地址
	var c *int
	fmt.Println(a)
	fmt.Println(b) // 读出指针指向的内存地址
	fmt.Println(*b) // 读出指针指向的内存地址，存的值
	fmt.Println(c) // nil
	fmt.Println(c == nil) // true

	a = 100 // 修改a
	println(a)
	println(*b)

	*b = 200 // 修改指针指向内存地址存的值
	println(a)
	println(*b)


	// demo2
	// i, j := 42, 2701
	// p := &i         // point to i
	// fmt.Println(*p) // read i through the pointer
	// *p = 21         // set i through the pointer
	// fmt.Println(i)  // see the new value of i
	// p = &j         // point to j
	// *p = *p / 37   // divide j through the pointer
	// fmt.Println(j) // see the new value of j


}
