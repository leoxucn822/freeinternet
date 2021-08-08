package main

import "fmt"

type calcSum []int

type sumResult interface {
	add(num1 int, num2 int)
}

func (s *calcSum) add(num1 int, num2 int) {
	fmt.Println(num1 + num2)
}

func main() {
	var num calcSum
	num.add(3, 5)
}
