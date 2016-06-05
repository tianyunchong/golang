package main

import (
	//"config" 本地调用
	"fmt"

	"github.com/tianyunchong/golang/src/config"
)

func main() {
	config.LoadConfig()
	fmt.Println("Hello, go!")
}
