/**
 * 测试切变的使用
 */
package main

import (
	"fmt"
)

//slice的声明
/*func main() {
	var s1 []int
	fmt.Println(s1)
}*/
//******************//
/*func main() {
	var a = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var s1 []int = a[5:10]
	fmt.Println(a)
	fmt.Println(s1)
}*/
//******************//
//slice使用make生命
/*func main() {
	s1 := make([]int, 3, 10) //第三个参数初始分配内存空间
}*/
//********************//
//append追加
/*func main() {
	var s1 = make([]int, 3, 6)
	fmt.Println(s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)
}*/
//**********************//
//copy
func main() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9}
	copy(s1, s2)
	fmt.Println(s1)
}
