// +build ignore

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	fmt.Println(pow)
	for i, v := range pow {
		fmt.Printf("index: %d,  value: %d\n", i, v)
	}

	// 效果一样
	for i :=0; i < len(pow); i++ {
		fmt.Printf("index: %d,  value: %d\n", i, pow[i])
	}
}


