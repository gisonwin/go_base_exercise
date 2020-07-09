package main

import (
	"fmt"
	"time"
)

/***
chan keyword 关键字管道 类似于linux pipe
后面跟着类型,第二个参数是大小
make是关键字,类似于new

*/
func testPipe() {
	fmt.Println(time.Now())
	pi := make(chan int, 3)
	pi <- 1
	pi <- 2
	pi <- 3
	t1 := <-pi
	print(t1)
	//fmt.Println(t1)
	pi <- 44
	//fmt.Println(len(pi))
	print(len(pi))
	fmt.Println(time.Now())

}

func Public() {
	println("public method invoked ")
}
