package main

import (
	"fmt"
	"time"
)

func testgorountine() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("hello world!\n")
		time.Sleep(time.Second)
	}
}

func main() {
	go testgorountine()

	for i := 1; i <= 10; i++ {
		fmt.Printf("hello golang!\n")
		time.Sleep(time.Second)
	}
}
