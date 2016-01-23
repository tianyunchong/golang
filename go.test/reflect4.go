/**
 * 反射进行方法的调用
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

func (u User) Hello(name string) {
    fmt.Println("Hello", name, ", my name is ", u.Name)
}

func main() {
    u := User{1, "ok", 13}
    v := reflect.ValueOf(u)
    mv := v.MethodByName("Hello")
    args := []reflect.Value{reflect.ValueOf("joe")}
    mv.Call(args)
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
