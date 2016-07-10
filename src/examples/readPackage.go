package main

import (
	"config"
	"fmt"
	//"github.com/tianyunchong/golang/src/config"
)

func main() {
	config.LoadConfig()
	fmt.Println("Hello, go!")
}
