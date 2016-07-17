package main

import (
	"time"
	"fmt"
)

func main() {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	td := tm.Format("2006-01-02")
	loc, _ := time.LoadLocation("Local")
	tm2,_ := time.ParseInLocation("2006-01-02", td, loc)
	fmt.Println(fmt.Sprintf("%d", tm2.Unix()))
}
