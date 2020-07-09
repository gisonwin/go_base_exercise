package main

import "fmt"

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/6/27 16:56
 */
//返回生成自然数序列的管道,从2开始 2,3,4......
//函数内部启动一个goroutine生产序列,返回对应的管道.
func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

//管道过滤器,删除能被素数整除的数
//内部驱动一个goroutine生产序列,返回过滤后序列对应的管道.
func PrimeFilter(in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}
/***
我们先是调用 GenerateNatural() 生成最原始的从2开始的自然数序列。然后开始一个100次迭代的
循环，希望生成100个素数。在每次循环迭代开始的时候，管道中的第一个数必定是素数，我们先读取
并打印这个素数。然后基于管道中剩余的数列，并以当前取出的素数为筛子过滤后面的素数。不同的素
数筛子对应的管道是串联在一起的。
 */
func main() {
	ch := GenerateNatural()
	for i := 0; i < 20; i++ {
		prime := <-ch //新出现的素数
		fmt.Printf("%v:%v\n", i+1, prime)
		ch = PrimeFilter(ch, prime) //基于新素数构造的过滤器
	}
}
