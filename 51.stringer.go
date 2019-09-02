// +build ignore

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type Self struct {
	Name string
}
func (s Self) String() string {
	return fmt.Sprintf("%T(%v)", s.Name, s.Name)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}

	b := Self{"longmore"}
	fmt.Println(a, z)

	fmt.Println(b)
}
