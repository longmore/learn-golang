// +build ignore

package main
import (
    "encoding/json"
    "fmt"
    "log"
)

type User struct {
    Name string
    Age int64
}


func main() {
  data := []byte(`{"Name":"酷走天涯","Age":27}`)
  var user *User
  error := json.Unmarshal(data, &user)
  if error != nil {
      log.Fatal(error)
  }
  fmt.Println(user)
}