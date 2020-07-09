package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/26 20:22
 */

//生产者
func Producer(factor int, out chan<- string) {
	for i := 0; ; i++ {
		out <- strings.Join([]string{strconv.Itoa(i * factor), strconv.Itoa(factor)}, "==>")
	}
}

//消费者
func Consumer(in <-chan string) {
	for v := range in {
		fmt.Println(v)
	}
}

/**
两个生产者,但这两个生产者之间并无同步事件可参考,它们是并发的,因此消费者输出的结果序列的顺序是不确定的,
生产者和消费者依然可以相互配合工作.
*/
func main() {
	ch := make(chan string, 64) //存储队列

	go Producer(3, ch) //生成3的倍数序列
	go Producer(5, ch) //生成5的倍数序列
	go Consumer(ch)    //消息生成的队列

	//Ctrl +C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v) \n", <-sig)
}
