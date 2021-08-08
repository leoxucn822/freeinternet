package main

import "fmt"

func main() {
	array := make([]int, 10)
	array = []int{9, 3, 4, 1, 13, 25, 10, 2, 99, 100}
	for i := range array {
		if array[i]%2 == 0 {
			fmt.Println(array[i])
		}
	}
}
