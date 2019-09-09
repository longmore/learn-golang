// +build ignore

package main

import (
    "os"
	"fmt"
)

func main() {
	//  ---------- 写文件 -----------
	dir, _ := os.Getwd()
	fmt.Println(dir)
	dir2 := fmt.Sprintf("%v/os", dir)
	fmt.Println(dir2)
	os.Chdir(dir2)

	// 打开文件, 替换文件内容写法
	// file, error := os.OpenFile("test.txt", os.O_RDWR, 0666)

	// 打开文件, 增加内容写法
	file, error := os.OpenFile("test.txt", os.O_RDWR|os.O_APPEND, 0666)

	//需要关闭打开的文件流
	defer file.Close()

	if error != nil {
		fmt.Println("error")
		fmt.Println(error)
	}
	_, error = file.WriteString("\n\t 写点东西进去...")
	fmt.Println(error)


	//  ---------- 读文件 -----------

	file2, error := os.Open("read.txt")
	defer file2.Close()

	fileInfo, _ := file2.Stat()
  	fmt.Println(fileInfo.Size()) // 单位为B

	// 创建缓冲区
	data := make([]byte,fileInfo.Size())

	_, error = file2.Read(data)

	fmt.Println(error)
	fmt.Println("读出来的内容为：")
 	fmt.Println(string(data))

}