package main

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/24 16:06
 */
var a string
var does = make(chan bool)
var msg string

func f() {
	print("a==>", a)
}
func hello() {
	a = "hello,world"
	print("hello a==>", a)
	go f()
}

func main() {
	//testOrderConsistency()
	//hello()
	go aGoroutine()
	<-does
}
func aGoroutine() {
	msg = "Gison Win"
	does <- true
	print(msg)
}

/***
使用chan 通道通信.
*/
func testOrderConsistency() {
	done := make(chan int)
	go func() {
		println("hello")
		done <- 1
	}()
	<-done
}
