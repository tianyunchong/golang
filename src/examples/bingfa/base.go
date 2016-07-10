// test to use channel
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(str string, ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("%s %d", str, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Println("you say: %q\n", <-c)
	}
	fmt.Println("you're boring! i'm leaving.")
}
