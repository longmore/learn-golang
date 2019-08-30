// +build ignore

// 函数闭包
package main

import "fmt"

func adder() func(int) int {

	sum := 0
	fmt.Println("sum: ")
	fmt.Println(sum)
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	a, b := adder(), adder()

	// fmt.Println(adder)
	fmt.Println(a(1))
	fmt.Println(a(2))
	fmt.Println(a(3))
	fmt.Println(b(100))
	fmt.Println(b(200))
	fmt.Println(b(300))
}