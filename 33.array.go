// +build ignore

package main

import "fmt"

func main() {
	var a[2]int
	a[0] = 44
	a[1] = 12
	fmt.Println(a)
	fmt.Println(a[0], a[1])

}