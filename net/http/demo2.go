// +build ignore

package main

import (
    "net/http"
    "fmt"
)

// 第一步 Handle 实现接口
type  Controller struct {
}

func (c Controller) ServeHTTP(resp http.ResponseWriter, req *http.Request){
    fmt.Println(req.Host)
    resp.Write([]byte("我是服务器返回的数据"))
    req.Body.Close()
}

func main() {
   // 第二步 监听服务
   http.ListenAndServe(":8081", &Controller{})
 }