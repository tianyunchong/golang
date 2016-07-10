/**
 * map的key, value交换实现
 */
package main

import (
	"fmt"
)

func main() {
	var m1 = map[int]string{1: "aa", 2: "bb", 3: "ccc"}
	var m2 = map[string]int{}
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)
}
