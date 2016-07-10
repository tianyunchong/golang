package main

import (
	"fmt"

	"github.com/tianyunchong/golang/test/src/config"
)

func main() {
	config.LoadConfig()
	fmt.Println("testaaa")
}
