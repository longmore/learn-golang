// +build ignore

package main

import (
  "time"
  "fmt"
)

func main() {
	// 1s 等于 1000ms
	second := time.Second
	fmt.Println(int64(second / time.Millisecond))

	// 10s
	seconds := 10
	fmt.Println(time.Duration(seconds) * time.Second)
}
