/**
 * 类型转换
 */
package main

import "fmt"

func main() {
	//var a float32 = 1.1
	//var b int = int(a)
	var a int = 65
	var b string = string(a)
	fmt.Println(b)
}
