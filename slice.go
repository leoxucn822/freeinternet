package main

import (
	"fmt"
)

func a() {
	x := []int{} //no length means this is slice, not array
	x = append(x, 0)
	x = append(x, 1)
	y := append(x, 2)
	z := append(x, 3)
	fmt.Println(y, z) // x = [0, 1],  y = [0, 1, 2], z = [0, 1, 2, 3]
}

func b() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	y := append(x, 3)
	z := append(x, 4)
	fmt.Println(y, z)
	fmt.Printf("%v\n", cap(x))
}

func main() {
	a()
	b()

	s := []int{1, 2, 3, 4, 5}
	//r := s[1:]
	r := append(s[:2], s[3:]...) //r = {1,2} + {4,5}
	//r := append(s, 6)
	t := append(r, 7, 8)
	fmt.Printf("%p %v\n", s, s)
	fmt.Printf("%p %v\n", r, r)
	fmt.Printf("%p %v %v\n", t, t, cap(t))

	s1 := []int{1, 2, 3}
	s2 := []int{7, 8, 9}

	copy(s1, s2)
	fmt.Printf("%p %v\n", s1, s1)
	fmt.Printf("%p %v\n", s2, s2)
}
