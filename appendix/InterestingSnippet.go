package main

import "fmt"

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/7/3 15:11
 */
func main() {
	//fmt.Printf("%s%c%s%c\n",q,0x60,q,0x60)
	//a, b := 2, 3
	//max := TernaryExpression(a > b, a, b).(int)
	//println(max)

	//ForrbiddenMainExit()

	RandomFake()
}

/**
自重写程序
Unix/Go语言之父产Ken Thompson 写过一个C语言的自重写程序
main(a){printf(a="main(a){printf(a=%c%s%c,34,a,34);}",34,a,34);}
我们提供一个Go语言版本的
*/
var q = `/* Go quine */
package main
import "fmt"
func main(){
	fmt.Printf("%s%c%s%c",q,0x60,q,0x60)
}
var q=`

/**
三元表达式
*/
func TernaryExpression(condition bool, trueVal, falseValue interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseValue
}

/**
禁止main函数退出的方法
*/
func ForrbiddenMainExit() {
	defer func() { select {} }()
	defer func() { <-make(chan bool) }()
}

/**
基于管道的随机数生成器
随机数的特点是不好预测,如果一个随机数的输出是可以简单预测的,一般会称为伪随机数.
*/
func RandomFake() {
	for i := range random(10) {
		fmt.Println(i)
	}
}
func random(n int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < n; i++ {
			select {
			case c <- 0:
			case c <- 1:
			case c <- 2:
			case c <- 3:
			case c <- 4:
			case c <- 5:
			case c <- 6:
			}
		}
	}()
	return c
}

