package main

import (
	"caiji/helper"
	"fmt"
)

func main() {
	mainConfig := helper.ReadIni("/www/gitwork/golang/src/caiji/config/config.ini", "main")
	fmt.Println(mainConfig)
}
