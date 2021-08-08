package main

import "fmt"

func main() {
	/*	var names[10] string

		fmt.Printf("%v\n", names)
		fmt.Printf("%v\n", len(names))*/

	arr1 := []int{1, 2, 3, 4}
	arr3 := [...]string{"hello", "world", "leo"}
	arr2 := arr1
	arr2[0] = 10
	//arr4 := make([]int, 5)
	arr4 := make([]int, 3, 100) //arr4 capacity is 100, but length is 3

	fmt.Printf("%v\n", arr1)
	fmt.Printf("%v\n", arr2)
	fmt.Printf("%v\n", arr3)
	fmt.Printf("%v %v\n", len(arr3), cap(arr3))
	arr4 = append(arr4, 5)
	arr4 = append(arr4, []int{6, 7, 8, 9, 10}...)
	fmt.Printf("%v\n", arr4)

	s := []int{1, 2, 3, 4}
	r := append(s[:2:2], s[3:]...)
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", r)
	fmt.Printf("%v\n", s)

	t := [...]int{1, 2, 3, 4}
	m := t
	fmt.Printf("%p %v\n", &t, t)
	fmt.Printf("%p %v\n", &m, m)
}
