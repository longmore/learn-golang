
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "reflect"
)

type User struct {
    name string
    age int64
}


func main() {
    user := User{"张三", 25}
    jsonStr, error := json.Marshal(user)
    if error != nil {
        log.Fatal(error)
    }
    fmt.Println(jsonStr)
    fmt.Println(string(jsonStr))
    fmt.Println(reflect.TypeOf(jsonStr))
    fmt.Println(reflect.TypeOf(string(jsonStr)))

    fmt.Println("--------------")

    data := []byte(`{"name":"李四","age":12}`)
    var user2 *User
    err := json.Unmarshal(data, &user2)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(user2)

}