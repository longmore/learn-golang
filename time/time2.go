// +build ignore

// 延迟d秒后返回通过信道返回时间值
package main

import (
	"time"
	"fmt"
)

func main() {
	date := <- time.After(3*time.Second)
	fmt.Println(date)
}