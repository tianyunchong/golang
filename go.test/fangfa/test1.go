/**
 * 测试方法
 */
package main

import (
    "fmt"
)

type A struct {
    Name string
}

type B struct {
    Name string
}

func main() {
    a := A{}
    a.Print()
}

func (a A) Print() {
    fmt.Println("A")
}
