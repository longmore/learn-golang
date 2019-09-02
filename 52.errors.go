// +build ignore

package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
	Time int
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s, %v", e.When, e.What, e.Time)
}

func run() error {
	return &MyError{ time.Now(), "it didn't work", 2, }
}

func main() {
	err := run();
	if err != nil {
		fmt.Println(err)
	}
}
