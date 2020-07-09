package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/24 16:56
 */
func main() {
	//testErrorMutex()
	//testFixedMutex()
	//testChannel()
	//testChannelWithCache()
	testCacheChannel(1)
	//testMultiThreadWithWaitGroup(10)

}

func testMultiThreadWithWaitGroup(size int) {
	var wg sync.WaitGroup
	for i:=0;i< size ;i++  {
		wg.Add(1)
		go func() {
			println("hello world ",i)
			wg.Done()
		}()
	}
	wg.Wait()

}
func testCacheChannel(size int)  {
	done:= make(chan int,size)
	//开启size个后台打印线程
	for i:=0;i<cap(done) ;i++  {
		//println(i)
		go func() {
			fmt.Printf("你好 世界 %d,now is %v\n",Goid(),time.Now())
			//time.Sleep(time.Second)
			done <- i
		}()
	}
	//等待后台线程完成
	for i:=0;i<cap(done) ;i++  {
		<- done
	}

}
/**
*对于带缓冲的channel,第K个接收完成操作发生在第K+C个发送操作完成之前,其中C是channel的缓存大小.虽然channel是带缓存的,main
线程接收完成是在后台线程发送开始但还未完成的时刻,此时打印工作已完成.
 */
func testChannelWithCache() {
	done := make(chan int, 1) //带缓存管理,缓存大小为1
	go func() {
		fmt.Println("你好 gisonwin")
		done <- 1
	}()
	<-done
}

/***
* 根据Go语言内存模型规范,对于从无缓冲channel进行的接收,发生在该channel进行的发送完成之前.所以main线程<-done会阻塞,等待
goroutine线程的done <- 1发送操作完成后,才能退出.而此时打印的工作已完成.注意这里main线程 <-done和goroutine里的 done<-1
可以互换位置.不影响结果输出,只是影响两个线程谁先退出的顺序.
*/
func testChannel() {
	var done = make(chan int)
	go func() {
		println("goid:",Goid())
		fmt.Println("gisonwin 你好")
		println(runtime.NumGoroutine())
		done <- 1
	}()
	<-done
}
func Goid() int {
	defer func()  {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)     }
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}


/**
* 修复的方式是在main函数所在线程中执行两次mu.Lock(),当第二次加锁时会因为锁已经被占用(不是递归锁)而阻塞,main函数的阻塞
状态驱动后台线程继续向前执行.当后台线程执行到mu.Unlock()时解锁,此时打印工作已完成,解锁会导致main函数中第二个mu.Lock()
阻塞状态取消,此时后台线程和主线程再没有其他的同步事件参考,它们退出的事件将是并发的:在main函数退出导致程序退出时,后台
线程可能已经退出了,也可能没有退出,虽然无法确定两个线程的退出时间,但是打印工作是可以正确完成的.

tips:使用sync.Mutex互斥锁同步是比较低级的做法.
*/
func testFixedMutex() {
	var mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Println("gisonwin 你好")
		mu.Unlock()
		fmt.Println("goroutine")
	}()
	mu.Lock()
	fmt.Println("main thread")
}

/**
*  互斥锁实现同步通信.这里不能直接对一个未加锁状态的sync.Mutex进行解锁,这会导致运行时异常.
因为mu.Lock,mu.Unlock()并不在一个Goroutine中,所以也就不满足顺序一致性内存模型.同时它们也没有其他的同步事件可以参考,这两
个事件不可排序也就是可以并发,由于是并发,所以main函数中的mu.Unlock()很有可能先发生,而这时mu互斥对象还处于未加锁状态,从而
导致运行时异常.
*/
func testErrorMutex() {
	var mu sync.Mutex
	go func() {
		fmt.Println("hello gisonwin")
		mu.Lock()
	}()
	mu.Unlock()
}
