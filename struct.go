package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Name string `json:"姓名"`
	Age  int    `json:"年龄"`
}

func main() {
	var e employee
	data, err := json.Marshal(e)
	f := new(employee)
	f.Name = "leo"
	f.Age = 18
	if err != nil {
		fmt.Println("json coding error!")
		return
	}
	fmt.Println(string(data))
	fmt.Println(f.Name)

	var a1 = employee{"xu leo", 18}
	fmt.Println(a1)
}
