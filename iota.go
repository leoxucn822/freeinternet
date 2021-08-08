package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

/*const (
	m = iota
)*/

const (
	n1 = iota //0
	n2 = 100  //100
	n3 = iota //2
	n4        //3

)
const (
	a1, b1 = iota + 1, iota + 2 //1,2
	c1, d1                      //2,3
	e1, f1                      //3,4
)

func main() {
	/*	var hodType int
		fmt.Printf("%v\n", hodType)
		fmt.Printf("%v\n", b)
		fmt.Printf("%v\n", hodType == b)*/
	//fmt.Printf("%v\n", b)

	fmt.Printf("%v\n", KB)
	fmt.Printf("%v\n", MB)
	fmt.Printf("%v\n", GB)

	fmt.Printf("%v\n", n1)
	fmt.Printf("%v\n", n2)
	fmt.Printf("%v\n", n3)
	fmt.Printf("%v\n", n4)
	fmt.Printf("%v %v %v %v %v %v\n", a1, b1, c1, d1, e1, f1)
}
