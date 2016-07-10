/**
 * 函数的定义使用
 */
package main

import (
	"fmt"
)

//**********************************
/*func main() {
	test(1, 2, 3)
}*/

/*
func test(a int, b int, c int) {

}
*/
//**********************************
/** 多个参数都是int型 */
/*
func test(a, b, c int) {

}
*/
/** 返回值 */
//**********************************
/*func test() {
	a, b, c = 1, 2, 3
	return a, b, c
}*/
/** 不定长变参，必需作为参数最后一个 */
/*func test(a ...int) {
	fmt.Println(a)
}*/
//**********************************
/** 值类型传递，改变内容 */
/*func main() {
	var a int = 4
	test(&a)
	fmt.Println(a)
}
func test(a *int) {
	*a = 3
	fmt.Println(*a)
}*/
//**********************************
/** 匿名函数 */
/*func main() {
	a := func() {
		fmt.Println("FuncA A")
	}
	a()
}*/
//**********************************
/** 闭包 返回一个匿名函数 */
/*func main() {
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}*/

//**********************************
/** defer 析构函数  即使遇到严重错误，也会执行, 逆序向上调用 */
/*func main() {
	fmt.Println("a")
	defer fmt.Println("b")
	defer fmt.Println("c")
}*/
//**********************************
/** defer的作用时调用某个函数, 循环后main return的时候才回调用i, 这时候已经是3了 */
/*func main() {
	for i := 1; i <= 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}*/

//**********************************
/** 触发异常panic */
func main() {
	A()
	B()
	C()
}
func A() {
	fmt.Println("Func A")
}
func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Rocover in B")
		}
	}()
	panic("Panic in B")
}
func C() {
	fmt.Println("Func C")
}
