/**
 * map功能测试学习
 */
package main

import (
    "fmt"
)

/*func main() {
	//var m map[int]string
	//m = map[int]string{}
	m := make(map[int]string)
	m[1] = "OK"
	a := m[1]
	fmt.Println(m)
	fmt.Println(a)
	delete(m, 1)
	fmt.Println(m)
}*/
//**********************//
//复杂的map
/*func main() {
	var m map[int]map[int]string
	m = make(map[int]map[int]string)
	m[1] = make(map[int]string)
	m[1][1] = "Ok"
	a := m[1][1]
	fmt.Println(a)
}*/
//********************//
//多返回值
/*func main() {
	var m map[int]map[int]string
	m = make(map[int]map[int]string)
	a, ok := m[2][1]
	fmt.Println(a, ok)
}*/
//********************//
//初始化赋值
func main() {
    var m = map[int]string{1: "test", 2: "bbb"}
    fmt.Println(m)
}
