/**
 * 反射测试学习
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

func (u User) Hello() {
    fmt.Println("hello wordld.")
}

func main() {
    u := User{1, "OK", 12}
    Info(u)
}

func Info(o interface{}) {
    t := reflect.TypeOf(o)
    fmt.Println("Type:", t.Name())
    v := reflect.ValueOf(o)
    fmt.Println("Fields:")
    for i := 0; i < t.NumField(); i++ {
        f := t.Field(i)
        val := v.Field(i).Interface()
        fmt.Println("%6s:%v = %v\n", f.Name, f.Type, val)
    }

    for i := 0; i < t.NumMethod(); i++ {
        m := t.Method(i)
        fmt.Println("%6s: %v\n", m.Name, m.Type)
    }
}
