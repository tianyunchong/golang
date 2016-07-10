package main


import "fmt"

func main() {
	arr := []interface{"aaa", "bbb"}
	test(arr...)
}


func test(a ...interface{}) {
	fmt.Println(a)
}

