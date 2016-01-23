/**
 * struct 结构测试使用
 */
package main

import (
	"fmt"
)

//**********************************
/** struct 的声明 */
/*
type test struct{}

func main() {
	a := test{}
	fmt.Println(a)
}
*/
//**********************************
/*
type Person struct {
	Name string
	Age  int
}

func main() {
	a := Person{}
	a.Name = "tianyunzi"
	a.Age = 27
	fmt.Println(a)
}
*/
//**********************************
/** 简便的使用方法 */
/*
type Person struct {
	Name string
	Age  int
}

func main() {
	a := Person{
		Name: "tianyunzi",
		Age:  27,
	}
	fmt.Println(a)
}
*/
//**********************************
/** 指针 传递参数 */
/*
type Person struct {
	Name string
	Age  int
}

func main() {
	a := Person{
		Name: "tianyunzi",
		Age:  27,
	}
	fmt.Println(a)
	A(&a)
	fmt.Println(a)
}

func A(a *Person) {
	a.Age = 20
	fmt.Println(a)
}
*/
//**********************************
/** 结构取地址初始化 */
/*
type Person struct {
	Name string
	Age  int
}

func main() {
	a := &Person{
		Name: "tianyunzi",
		Age:  27,
	}
	fmt.Println(a)
	A(a)
	fmt.Println(a)
}

func A(a *Person) {
	a.Age = 20
	fmt.Println(a)
}
*/
//**********************************
/** 匿名结构 */
/*
func main() {
	a := struct {
		Name string
		Age  int
	}{
		Name: "tianyunzi",
		Age:  27,
	}
	fmt.Println(a)
}
*/
//**********************************
/** 匿名字段 */
/*
type Person struct {
	string
	int
}

func main() {
	a := Person{"tianyunzi", 19}
	fmt.Println(a)
}
*/
//**********************************
/** 共同属性, 嵌入结构 */
type human struct {
	Sex int
}
type teacher struct {
	human
	Name string
	Age  int
}
type student struct {
	human
	Name string
	Age  int
}

func main() {
	a := teacher{Name: "tianyunzi", Age: 19, human: human{Sex: 0}}
	b := student{Name: "tianyunchong", Age: 21, human: human{Sex: 1}}
	a.Age = 13
	b.Sex = 100
	fmt.Println(a, b)
}
