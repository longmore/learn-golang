// +build ignore

package main

import (
	"time"
	"fmt"
	"reflect"
)

func main() {
	d := time.Now()
	fmt.Println("打印出当前时间")
	fmt.Println(d)

	year, month, day := time.Now().Date()
	fmt.Println("年月日获取")
	fmt.Printf("%v %v %v \n", year, month, day)

	fmt.Println(reflect.TypeOf(month))
	fmt.Println(reflect.TypeOf(month.String()))

	fmt.Println("---section 1----")

	fmt.Println("月份值判断")
	if month == time.September {
		fmt.Println("当前月份是 9月")
	}

	fmt.Println("---section 2----")

	// "2006-01-02 15:04:05" 必须不能变，golang指定的时间原点
	d1 := d.Format("2006-01-02 15:04:05")
	fmt.Println("输出当前时间，并格式化")
	fmt.Println(d1)

	d2, _ := time.Parse("2006-01-02 15:04:05", "2019-06-15 08:37:18")
	fmt.Println("输出当前时间，并格式化")
	fmt.Println(d2)

	fmt.Println("---section 3----")

	d3 := d.Unix()
	fmt.Println("时间戳")
	fmt.Println(d3)

	fmt.Print(d.Year())
	fmt.Println("年")

	fmt.Print(d.Month())
	fmt.Println("月")

	fmt.Print(d.Day())
	fmt.Println("日")

	fmt.Print(d.Hour())
	fmt.Println("时")

	fmt.Print(d.Minute())
	fmt.Println("分")

	fmt.Print(d.Second())
	fmt.Println("秒")

	fmt.Print(d.Weekday())
	fmt.Println("星期")

	fmt.Println("---section 4----")

	dd := time.Date(2000, 1, 02, 12, 00, 00, 0, time.UTC)
	fmt.Println(dd)

	fmt.Println("时间对比")
	fmt.Println(d.After(dd))
	fmt.Println(d.Before(dd))
	fmt.Println(d.After(d))

	ddd := time.Date(2020, 9, 02, 12, 00, 00, 0, time.UTC)
	fmt.Println(ddd)

	fmt.Println("---section 5----")

	fmt.Println("时间差")
	diff1 := d.Sub(dd)
	diff2 := d.Sub(ddd)
	fmt.Println(diff1)
	fmt.Println(diff2)

	fmt.Println(diff1.Hours())

	fmt.Println("---section 6----")

	fmt.Println("时间差值")
	// 1分钟前
	dur1, _ := time.ParseDuration("-1m")
	fmt.Println(dur1)
	// 1小时前
	dur2, _ := time.ParseDuration("-1h")
	fmt.Println(dur2)
	// 1天后
	dur3, _ := time.ParseDuration("24h")
	fmt.Println(dur3)

	fmt.Println("时间相加")

	fmt.Println("现在")
	fmt.Println(d.Format("2006-01-02 15:04:05"))
	fmt.Println("1分钟前")
	add1 := d.Add(dur1)
	fmt.Println(add1.Format("2006-01-02 15:04:05"))
	fmt.Println("1小时前")
	add2 := d.Add(dur2)
	fmt.Println(add2.Format("2006-01-02 15:04:05"))
	fmt.Println("1天后")
	add3 := d.Add(dur3)
	fmt.Println(add3.Format("2006-01-02 15:04:05"))

}

// https://blog.csdn.net/u011525168/article/details/88401690