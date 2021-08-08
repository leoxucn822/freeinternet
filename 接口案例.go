package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type student1 struct {
	name  string
	age   int
	score int
}

type studentslice []student1

func (stu studentslice) Len() int {
	return len(stu)
}

func (stu studentslice) Less(i, j int) bool {
	return stu[i].score < stu[j].score
}

func (stu studentslice) Swap(i, j int) {
	stu[i].score, stu[j].score = stu[j].score, stu[i].score
}

func main() {
	var stu studentslice
	for i := 0; i < 10; i++ {
		s := student1{
			name:  fmt.Sprintf("英雄%d", rand.Intn(100)),
			age:   rand.Intn(100),
			score: rand.Intn(100),
		}
		stu = append(stu, s)
	}
	for _, v := range stu {
		fmt.Println(v)
	}
	fmt.Printf("\n")
	//排序之后的stu
	sort.Sort(stu)
	for _, v := range stu {
		fmt.Println(v)
	}
}
