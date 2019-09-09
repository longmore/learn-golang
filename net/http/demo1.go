package main

import (
    "net/http"
    "log"
    "io/ioutil"
    "fmt"
)

func main() {
  	// 创建客户端发get请求
	res, error := http.DefaultClient.Get("https://m.acfun.cn/app/download")

	/**
		type Response struct {
			Status     string // e.g. "200 OK"
			StatusCode int    // e.g. 200
			Proto      string // e.g. "HTTP/1.0"
			ProtoMajor int    // e.g. 1
			ProtoMinor int    // e.g. 0
			Header Header  // 相应头部
			Body io.ReadCloser
			ContentLength int64 // 内容长度
			TransferEncoding []string // 传输数据编码
			Close bool // 是否关闭连接
				Uncompressed bool //是否未压缩
			Trailer Header
			Request *Request // 原始请求相关
			TLS *tls.ConnectionState // https 加密相关
		}
	**/

	// 相应结束后,请及时关闭会话
	if error != nil {
		fmt.Println(error)
	}
	defer res.Body.Close()
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("cookie值：")
	fmt.Println(res.Cookies())
	fmt.Println("--------------")
	fmt.Println("Header值：")
	fmt.Println(res.Header["Content-Type"])
	fmt.Println("--------------")
	// 读取相应的数据
   	data, error := ioutil.ReadAll (res.Body)
   	fmt.Println(string(data))
}

// https://blog.csdn.net/u011525168/article/details/88401872