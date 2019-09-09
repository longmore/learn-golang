// +build ignore

package main

import (
	"fmt"
)

type USB interface {
	start()
	end()
}

type Computer struct {
	name string
}

func (c Computer) start() {
	fmt.Println(c.name, "start")
}

func (c Computer) end() {
	fmt.Println(c.name, "end")
}

type Phone struct {
	name string
}

func (ph Phone) start() {
	fmt.Println(ph.name, "start")
}

func (ph Phone) end() {
	fmt.Println(ph.name, "end")
}

func main() {
	var com Computer = Computer{"三星"}
	Option(com)

	var ph Phone = Phone{"苹果手机"}
	Option(ph)

}

func Option(in USB){
	in.start()
	in.end()
}