package main

import (
	"net/http"
	"fmt"
)

func main()  {
	GetCode("http://www.baidu.com")
}

func GetCode(url string) {
	res, err := http.Head(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.StatusCode)
}