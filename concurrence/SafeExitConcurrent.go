package main

import (
	"fmt"
	"sync"
	"time"
)

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/27 17:32
 */
/***
有时我们需要通知goroutine停止它正在干的事情,特别是它工作在错误的方向上的时候.Go语言并没有提供在一个直接终止goroutine的方法
,由于这样会导致goroutine之间共享变量处在未定义状态.但如果我们想要退出两个或多个goroutine怎么办呢?
Go语言中不同goroutine之间主要依靠管道进行通信和同步.要同时处理多个管理的发送或接收操作,我们需要使用select关键字.当select
有多个分支时,会随机选择一个可用的管道分支,如果没有可用的分支则选择default分支,否则会一直保存阻塞状态.基于select实现的管道
超时判断
select {
	case v:= <- in:
		fmt.Println(v)
	case <-time.After(time.Second):
		return //超时
}
通过select的default分支实现非阻塞的管道发送或接收操作:
select{
	case v:= <-in:
		fmt.Println(v)
	default:
		//没有数据
}
通过select来阻止main函数退出
func main(){
	//do some things
	select{}
}
*/
func testRandom() {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			}
		}
	}()
	for v := range ch {
		fmt.Println(v)
	}
}
func main() {
	//testRandom()//select 生成随机数
	//testExitControl()
	//testExitControl2()
	testExitControl3()
}
func worker(cancel chan bool) {
	for {
		select {
		default:
			fmt.Println("hello") //工作正常
		case <-cancel:
			//退出
		}
	}
}
func testExitControl() {
	cancel := make(chan bool)
	go worker(cancel)
	time.Sleep(time.Second)
	cancel <- true
}

/***
管道的发送操作和接收操作是一一对应的,如果要停止多个goroutine那么可能需要创建同样数量的管道,这个代价太大了.我们可以通过
close()关闭一个管道来实现广播的效果,所有从关闭管道接收的操作均会收到一个零值和一个可选的失败标志.
*/
func worker2(cancel chan bool) {
	for {
		select {
		default:
			fmt.Println("worker2")
		//normal work
		case <-cancel:
			//exit
		}
	}
}

/***
我们通过close来关闭cancel管道向多个goroutine广播退出的指令.不过这个程序依然不够稳健:当每个goroutine收到退出指令退出时
一般会进行一定的清理工作,但是退出的清理工作并不能保证被完成,因为main线程并没有等待各个工作goroutine退出工作完成的机制.
我们可以结合sync.WaitGroup来改进.
*/
func testExitControl2() {
	cancel := make(chan bool)
	for i := 0; i < 10; i++ {
		go worker2(cancel)
	}
	time.Sleep(time.Second)
	close(cancel)
}
func worker3(wg *sync.WaitGroup, cancel chan bool) {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("worker3")
		case <-cancel:
			return
		}
	}
}
func testExitControl3() {
	cancel := make(chan bool)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker3(&wg, cancel)
	}
	time.Sleep(time.Second)
	close(cancel)
	wg.Wait()
}
