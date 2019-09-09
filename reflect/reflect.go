package main

import (
    "fmt"
    "reflect"
)

func main() {
    i := 10
	fmt.Printf("%d %T \n", i, i)
	fmt.Printf("%v %v \n", reflect.TypeOf(i), reflect.ValueOf(i))
	fmt.Printf("%v %v \n", reflect.TypeOf(i).Kind(), reflect.ValueOf(i))


	j := "中国人"
	fmt.Printf("%s %T \n", j, j)
	fmt.Printf("%v %v \n", reflect.TypeOf(j), reflect.ValueOf(j))
    fmt.Printf("%v %v \n", reflect.TypeOf(j).Kind(), reflect.ValueOf(j))
}