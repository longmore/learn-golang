package main

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
    u, err := url.Parse("http://www.baidu.com/search/result?name=joe&age=21")
    if err != nil {
        log.Fatal()
    }
    fmt.Println(u.Scheme) // http
    fmt.Println(u.Host)      // www.baidu.com
    fmt.Println(u.Path)      // /search
    fmt.Println(u.Query()) // map[age:[21] name:[joe]]
    fmt.Println(u.User)       // nil
    fmt.Println(u.RawQuery) // name=joe&age=21
    fmt.Println(u.Fragment)  // nil
    fmt.Println(u.Opaque)      // nil
    fmt.Println(u.Hostname())    // www.baidu.com
    fmt.Println(u.EscapedPath())  // /search/result
    fmt.Println(u.IsAbs())      // true
    fmt.Println(u.Port())        // nil
	fmt.Println(u.RawPath)  // nil

	fmt.Println("----- section 1 ------")

	fmt.Println("query encode: ")
	fmt.Println(url.QueryEscape("name=124&sex=123"))
	fmt.Println("query decode: ")
	fmt.Println(url.QueryUnescape("name%3D124%26sex%3D123"))


	fmt.Println("----- section 2 ------")

	u2, err := url.ParseRequestURI("http://bing.com/search?q=dotnet#name")
    if err != nil {
        log.Fatal()
    }
	fmt.Println(u2.String())
	fmt.Println(u2.RequestURI())

	fmt.Println("----- section 3 ------")

	// u3, err := url.Parse("http://www.baidu.com/search/result?name=joe&age=21")
	v := url.Values{}
    v.Set("name", "张三")
    v.Add("age", "21")
    v.Add("like", "apple")
    v.Add("like", "banana")
	fmt.Printf("name value: %v \n", v.Get("name"))
	fmt.Printf("like value: %v \n", v.Get("like"))
    fmt.Println(v.Encode())
    v.Del("name")
    fmt.Println(v.Encode())


}