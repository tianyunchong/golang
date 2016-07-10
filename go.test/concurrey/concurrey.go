/**
 * 并发测试使用
 */
package main

import (
    "fmt"
)

func main() {
    c := make(chan bool)
    go func() {
        //通知mian函数本函数已经执行完毕
        fmt.Println("Go Go Go!!!")
        c <- true
    }()
    <-c
}
