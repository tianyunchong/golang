/**
*  本案例可以看到channel默认是不会阻塞住进程的
*/
package main

import "fmt"

func test(c chan int) chan int{
	c <- 2
	c <- 1
	return c
}

func main() {
	c := make(chan int, 2)
	cs := test(c)
	fmt.Println(<-cs)
	fmt.Println(<-cs)
	fmt.Println("testaaa")
}
