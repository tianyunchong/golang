/**
 * 控制语句测试
 */
package main

import (
	"fmt"
)

//if语句的应用
/*func main() {
	var a int = 10
	if a := 1; a >= 1 {
		fmt.Println(a)
	}
	fmt.Println(a)
}*/
//****************************//
//循环语句, 支持3种形式
//形式1，无限循环
/*func main() {
	var a int = 10
	for {
		a++
		if a > 15 {
			break
		}
		fmt.Println(a)
	}
}*/
//形式2:自带表达式
/*func main() {
	var a int = 10
	for a < 15 {
		a++
		fmt.Println(a)
	}
}*/
//形势3:最常见的
/*func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}
	fmt.Println("over")
}*/
//******************************//
//switch语句
/*func main() {
	a := 1
	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}
}*/
func main() {
	a := 1
	switch {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough
	case a >= 1:
		fmt.Println("a>=1")
	default:
		fmt.Println("None")
	}
}
