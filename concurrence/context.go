package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/** context包用来简化对于处理单个请求的多个goroutine之间与请求域的数据,超时和退出等操作.我们可以用context包来重新实
现前面的线程安全退出或超时的控制.
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/27 18:54
*/
func worker4(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()
	for {
		select {
		default:
			fmt.Println("worker4")
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
/***
当并发体超时或main主动停止工作线程goroutine时,每个工作者都可以安全退出.
Go语言是带内存自动回收特性的,因此内存一般不会泄漏.
Go语言中大部分函数的代码结构几乎相同,首先是一系列的初始检查,用于防止错误发生,之后是函数的实际逻辑.
 */
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker4(ctx, &wg)
	}
	time.Sleep(time.Second)
	cancel()
	wg.Wait()
}
