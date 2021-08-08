package main

import (
	"fmt"
)

func main() {
	a := "hello world北京"
	b := []byte(a)
	c := []rune(a)
	fmt.Printf("%v\n", a)
	for i := range b {
		fmt.Printf("%v", string(b[i]))
	}
	fmt.Println()
	for i := range c {
		fmt.Printf("%v", string(c[i]))
	}
	fmt.Println()
	var ptr *string
	ptr = &a
	fmt.Printf("%p\n", ptr)

	var e = [...]int{1, 2, 3, 4, 5}
	var f = make([]string, 0)
	var g [5]int
	var number int
	fmt.Scanln(&number)
	var h = make([]string, number)

	fmt.Println(e)
	for i := 0; i < 2; i++ {
		fmt.Scanf("%s", &h[i])
		fmt.Println(h[i])
	}
	f = append(f, h...)
	//f = append(f, h[2])
	fmt.Println(f)
	fmt.Println(g)
}
