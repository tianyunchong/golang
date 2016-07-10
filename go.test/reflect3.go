/**
 * 测试使用结构体反射
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

func main() {
    u := User{1, "ok", 13}
    Set(&u)
    fmt.Println(u)
}

func Set(o interface{}) {
    v := reflect.ValueOf(o)
    //v.Elem().Canset是否可以被修改
    if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
        fmt.Println("XXX")
        return
    } else {
        v = v.Elem()
    }
    f := v.FieldByName("Name")
    if !f.IsValid() {
        fmt.Println("bad")
        return
    }
    if f.Kind() == reflect.String {
        f.SetString("BYEBYE")
    }
}
