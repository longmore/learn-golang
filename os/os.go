package main

import (
	"os"
	"fmt"
)

func main() {
	// 获取当前目录
	dir, _ := os.Getwd()
	fmt.Println(dir)

	// 查看所有环境变量
	envs := os.Environ()
	fmt.Println(envs)

	// 获取环境变量
	goroot := os.Getenv("GOROOT")
	gopath := os.Getenv("GOPATH")
	fmt.Println(goroot)
	fmt.Println(gopath)
}
// https://blog.csdn.net/u011525168/article/details/88401874