/**
 * 反射测试1
 */
package main

import (
    "fmt"
    "reflect"
)

type User struct {
    Id   int
    Name string
    Age  int
}

type Manage struct {
    User
    title string
}

func main() {
    m := Manage{User: User{1, "Ok", 12}, title: "123"}
    t := reflect.TypeOf(m)
    //fmt.Printf("%#v\n", t.Field(1))
    fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
}
