/**
 * 迭代操作
 */
package main

import (
	"fmt"
)

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	s1 := arr[:]
	for i, v := range s1 {
		fmt.Println(i, v)
	}
}
