package main

import (
	"fmt"
	"runtime"
	"time"
)

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2019/8/14 9:13
 */
/***
* goroutine是Go并行设计的核心.其实就是线程但它比线程更小,十几个gorougine可能体现在底层就是5,6个线程,Go语言内部
帮你实现了goroutine之间的内存共享.执行goroutine只需要极少的栈内存(大概4,5KB)当然会根据相应的数据伸缩.因此可同时
支行成千上万个并发任务.goroutine比thread更易用,高效,轻便.
goroutine是通过Go的runtime管理的一个线程管理器.通过go关键字实现了,就是一个普通函数.
通过关键字go就启动了一个goroutine.
*/

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() //表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine。
		fmt.Println(s)
	}
}

/***
遵循的原则:
            不要通过共享来通信，而要通过通信来共享
默认情况下，调度器仅使用单线程，也就是说只实现了并发。想要发挥多核处理器的并行，需要在我们的程序中显
示的调用 runtime.GOMAXPROCS(n) 告诉调度器同时使用多个线程。GOMAXPROCS 设置了同时运行逻辑代码的系统线
程的最大数量，并返回之前的设置。如果n < 1，不会改变当前设置。以后Go的新版本中调度得到改进后，这将被移除。
*/
func go_routine() {
	go say("world")
	say("hello")
}

func main() {
	//go_routine()
	//nobuffer_channel()
	//buffer_channel()
	//range_go()
	//select_go()
	select_timeout()
}

/***
默认情况下,channel接收和发送数据都是阻塞的,除非另一端已经准备好,这样就使得goroutine同步就的更加简单
而不需要显式lock.所谓阻塞就是如果读取value:= <-ch它将会被阻塞,直到有数据接收.其次,任何发送ch<-5都将
会被阻塞,直到数据被读出.无缓冲channel是在多个goroutine之间同步很棒的工具.
*/
func nobuffer_channel() {
	a := []int{7, 4, -3, 9, 8, 0, 2, 5}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

/***
channels:
goroutine支行在相同的地址空间,因此访问共享内存必须做好同步.goroutine之间如何进行数据通信呢,Go提供了一个很好的通信机制channel.
channel可以与Unix shell中的双向管道做类比:可以通过它发送或接收值.这些值只能是特定类型:channel类型.
定义一个channel时需要定义发送到channel的值的类型.注意必须使用make 创建channel.
ci:= make(chan int)
cs := make(chan string)
cf := make(chan interface{})
channel通过操作符<-来接收和发送数据
ch <- v //发送v到channel ch
v := <-ch //从ch中接收数据,并赋值给v


*/
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum //send sum to channel c
}

/***
Go也允许指定channel缓冲大小,就是channel可以存储多少元素.
ch := make(chan bool,4).创建可以存储4个bool型channel.在这个channel中,前4个元素可以无阻塞的写入.
当写入第5个元素时,代码将会被阻塞,直到其他goroutine从channel中读取一些元素,腾出空间.
ch := make(chan type,value)
value ==0 !无缓冲(阻塞)
value >0 ! 缓冲(非阻塞,直到value个元素)
*/
func buffer_channel() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

/****
Range and Close:
在上述例子中,我们需要读取两次c,Go可以通过range,像操作slice或者map一样操作缓存类型的channel.
*/
func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

/**
for i := range c能够不断的读取channel里面的数据，直到该channel被显式的关闭。上面代码我们看到可以显 式的关闭channel，
生产者通过关键字close函数关闭channel。关闭channel之后就无法再发送任何数据了，在消费 方可以通过语法v, ok := <-ch测试channel是否被关闭。
如果ok返回false，那么说明channel已经没有任何数据 并且已经被关闭。 记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，
这样容易引起panic 另外记住一点的就是channel不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显
式的结束range循环之类的
*/
func range_go() {
	c := make(chan int, 40)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

/***
如果存在多少channel的时候 ,我们可以通过关键字select,通过select可以监听channel上的数据流动.
select默认是阻塞的,只有当监听channel中有发送或接收可以进行时才会支行,当多个channel都准备好的时候
select是随机选择一个执行的.
*/
func fibonacci2(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, y+x
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

/***
在select里面还有default语法，select其实就是类似switch的功能，default就是当监听的channel都没有准备好
的时候，默认执行的（select不再阻塞等待channel）。
*/
func select_go() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci2(c, quit)
}

/***
有时候会出现goroutine阻塞的情况,我们如何避免整个程序进入阻塞情况呢.可以利用select来设置超时.
runtime goroutine
  Goexit:
  退出当前执行的goroutine,但是defer函数还是会继续调用
  Gosched
  让出当前goroutine的执行权限,调度器安排其他等等任务支行,并在下次某个时刻从该位置恢复执行
  NumCPU
  返回CPU核数量
  NumGoroutine
  返回正在执行和排队的任务总数
  GOMAXPROCS
   用来设置可以运行的CPU核数
*/
func select_timeout() {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				fmt.Println(v)

			case <-time.After(5 * time.Second):
				fmt.Println("timeout")
				o <- true
				break
			}
		}
	}()
	<-o
}
