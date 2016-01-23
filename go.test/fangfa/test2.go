/**
 * go方法指针传递
 *
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
    fmt.Println(a.Name)
}

func (a *A) Print() {
    a.Name = "AA"
    fmt.Println("A")
}
