package main

import (
	"time"
	"fmt"
)

func main() {
	d := time.Now()
	fmt.Println("打印出当前时间")
	fmt.Println(d)

	year, month, day := time.Now().Date()
	fmt.Println("年月日获取")
	fmt.Printf("%v %v %v \n", year, month, day)


	fmt.Println("月份值判断")
	if month == time.September {
		fmt.Println("当前月份是 9月")
	}

	// "2006-01-02 15:04:05" 必须不能变，golang指定的时间原点
	d1 := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("输出当前时间，并格式化")
	fmt.Println(d1)
}
