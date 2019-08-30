// +build ignore

package main

import "fmt"

func main() {
	pow := make([]int, 10)

	printSlice("pow", pow)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	printSlice("pow", pow)

	for i, value := range pow {
		fmt.Printf("%d, %d\n", i, value)
	}

	// 如果没有使用i，用_代替
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}