package main

import (
	//"time"
	//"fmt"
	"time"
)

func main() {
	//fmt.Print("Hello", "Gison",time.Now())
	/**
	验证goroute 协程并发,我们用时间来证明这300个确实是同时并发的.
	基于csp communicate sequential process模型实现
	 */
/*	for i:=0;i<300 ;i++  {
		go fmt.Println(i,time.Now())
	}
	time.Sleep(time.Second)*/
	for i:=0;i<10 ;i++  {
		go testPipe()
	}
	time.Sleep(time.Second)
	Public()
}
