package main

import (
	"fmt"
	"time"
)

func putData(intChan chan int) {
	for i := 2; i <= 1000000; i++ {
		intChan <- i
	}
	//关闭channel
	close(intChan)
}

func checkPrime(intChan chan int, primeChan chan int, exitChan chan bool) {
	//var num int
	var flag bool
	for {
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		// 判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			primeChan <- num
		}
	}
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 500000)
	exitChan := make(chan bool, 8)

	start := time.Now().Unix()
	// 开启一个协程，写入8000个数据
	go putData(intChan)

	// 开启四个协程，从intChan取出数据并判断是否是素数，如果是，就放在primeChan
	for i := 0; i < 8; i++ {
		go checkPrime(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(primeChan)
		end := time.Now().Unix()
		fmt.Println(end - start)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println(res)
	}
}
