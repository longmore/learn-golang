// +build ignore

package main

import (
    "time"
    "fmt"
)

func main() {
    c := time.Tick(1 * time.Second)
    for now := range c {
        fmt.Printf("延时执行该做的事情 %s\n", now)
    }

}