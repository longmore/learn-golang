// +build ignore

package main
import "fmt"

func main() {
	var a []int
	printSlice("a1", a)

	a = append(a, 12)
	printSlice("a2", a)

	a = append(a, 21)
	printSlice("a3", a)

	a = append(a, 31, 41, 51)
	printSlice("a4", a)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}