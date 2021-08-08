package main

import (
	"fmt"
	"sort"
)

func main() {

	var map1 map[int]string
	var map2 = make(map[int]string)
	var map3 = map[int]string{1: "go", 2: "C", 3: "C++", 4: "Java"}

	map3 = make(map[int]string)

	//map1[0] = "hello" // if map is nil, can't assignment
	fmt.Printf("%v\n", map1)
	fmt.Printf("%v\n", map2)
	fmt.Printf("%v\n", map3)

	fmt.Println(map1 == nil)
	fmt.Printf("%v\n", map3[40]) //print the none

	v, k := map3[40]

	if k {
		fmt.Printf("对应的数值是: %v\n", v)
	} else {
		fmt.Printf("操作的值不存在，获取的值是零值 %v\n", v)
	}

	var keys []int
	for k, _ := range map3 {
		keys = append(keys, k)
	}
	fmt.Println(keys) //打印出来是无序的
	sort.Ints(keys)
	fmt.Println(keys)

	monster := make([]map[string]string, 3)
	if monster[0] == nil {
		monster[0] = make(map[string]string, 2)
		monster[0]["Name"] = "老大"
		monster[0]["Age"] = "18"
	}
	if monster[1] == nil {
		monster[1] = make(map[string]string)
		monster[1]["Name"] = "牛魔王"
		monster[1]["Age"] = "20"
	}
	if monster[2] == nil {
		monster[2] = make(map[string]string)
		monster[2]["Name"] = "红孩儿"
		monster[2]["Age"] = "5"
	}
	newmonster := map[string]string{
		"Name": "张三",
		"Age":  "12",
	}
	monster = append(monster, newmonster)
	fmt.Println(monster)
}
