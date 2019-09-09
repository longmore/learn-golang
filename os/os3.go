// +build ignore

package main

import (
    "os"
  "fmt"
)

func main() {
	// 创建一个文件夹
	file,error:= os.Open("./os/test.txt")
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(file)
}