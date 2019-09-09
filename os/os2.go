// +build ignore

package main

import (
    "os"
  	"fmt"
)

func main() {
  	// 进入桌面目录
  	// 创建一个文件夹
	file,error := os.Create("./os/createFile.txt")
	if error != nil{
		fmt.Println("error")
		fmt.Println(error)
	}
 	fmt.Println(file)
}