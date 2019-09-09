// +build ignore

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	cpuNum := runtime.NumCPU()
	fmt.Println("当前系统CPU核数量:", cpuNum)
    fmt.Println("GO执行文件根目录:", runtime.GOROOT())
	fmt.Println("操作系统:", runtime.GOOS)

	fmt.Printf("设置几个核执行任务: %v \n", 1)
	runtime.GOMAXPROCS(1)

	t1 := time.Now().UnixNano()

	sum := 1
	for i := 1; i < 1000000; i++ {
		sum = sum * i
	}
	t2 := time.Now().UnixNano()
	diff := t2 - t1
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(sum)
	fmt.Println("运行结束")
	fmt.Println(diff)

}
