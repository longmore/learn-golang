// +build ignore

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	m["Age"] = 25
	m["Age2"] = 0
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	fmt.Println(m)
	// 通过双赋值检测某个键存在：
	v1, ok1 := m["Answer"]
	fmt.Println("The value:", v1, "Present?", ok1)
	v2, ok2 := m["Age"]
	fmt.Println("The value:", v2, "Present?", ok2)
	v3, ok3 := m["Age2"]
	fmt.Println("The value:", v3, "Present?", ok3)
}
