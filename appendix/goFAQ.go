package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"strconv"
)

/**当参数的可变参数是空接口类型时,传入空接口的切片时需要注意参数展开的问题.
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2020/7/3 9:22
 */
func VarParam() {
	var a = []interface{}{1, 2, 3}
	fmt.Println(a)
	fmt.Println(a...)
}

/**
函数调用中,数组是值传递,无法通过修改数组类型的参数返回结果.必要时需要使用切片slice.
*/
func ArrayTransform() {
	x := [3]int{1, 2, 3}
	func(arr [3] int) {
		arr[0] = 7
		fmt.Println("arr is: ", arr)
	}(x)
	fmt.Println("x is : ", x)
}

/***
map是一种hash表实现,每次遍历的顺序都可能不一样.
*/
func MapIterator() {
	m := map[string]string{
		"1": "one",
		"2": "two",
		"3": "three",
	}
	for key, value := range m {
		fmt.Println("k,v==>", key, ",", value)
	}
}

/**
在局部作用域中,命名的返回值内同名的局部变量屏蔽.
*/
func ReturnValueMask(err error) {
	if err != nil {
		fmt.Println("err not nil,is: ", err)
		return
	}
	fmt.Println("err is nil")
	return
}

/**
recover捕获的祖父级调用时的异常,直接调用时无效.
*/
func RecoverMustBePerformInDeferFunction() {
	/**
		all is invalid,such as :
	/////////////////////////////////
		recover()
		panic(1)
	//////////////////////////////////
		defer recover()
		panic(1)
	//////////////////////////////////
	多层嵌套依然无效
	defer func(){
		func(){recover()}()
	}()
	panic(1)
	*/
	//必须在defer函数中直接调用才有效
	defer func() {
		recover()
	}()
	panic(1)
}

/***
后台goroutine无法保证完成任务,
*/
func MainEarlyTermination() {
	println("hello")
}

/**
Goroutine是协作式抢占调度,它本身不会主动放弃CPU.下面例子go func 有可能会被打印,大部分时间不会被打印.
*/
func MonopolizeCPU() {
	runtime.GOMAXPROCS(1) //
	go func() {
		for i := 0; i < 5; i++ {
			println("another goroutine ==> ", i)
		}
		os.Exit(0)
	}()

	go func() { //
		for i := 0; i < 10; i++ {
			println("~~~~~~~~~~~~~~~~~~~~~", i)
		}
		//os.Exit(0)
	}()

	//for i:=0;i<2 ;i++  {
	//	println("main thread i==>",i)
	//	runtime.Gosched()
	//}
	select {}

}

var msg string
var done = make(chan bool)

func setup() {
	msg = "hello"
	done <- true
}

/**
闭包错误引用同一个变量
*/

func ClosureError() {
	for i := 0; i < 5; i++ {
		i := i
		println("i : ", i)
		defer func() {
			println(i)
		}()
	}
}

/**
defer在函数退出时才能执行,在fro 中执行defer会导致资源延迟释放.
解决办法 在for中构造一个局部函数,在局部函数中执行defer.
*/
func ClosureRight() {
	for i := 0; i < 5; i++ {
		defer func(k int) {
			println(k)
		}(i)
	}
}

func SliceLockArray() {
	headerMap := make(map[string][]byte)
	dir, _ := os.Getwd()
	fmt.Println("current path:", dir)
	path.Join(dir, )
	chdir := os.Chdir(dir)
	fmt.Println("chdir:", chdir)
	name := path.Join(dir, "/grpc/hello.pb.go")
	fmt.Println("name:", name)
	for i := 0; i < 5; i++ {
		data, err := ioutil.ReadFile(name)
		if err != nil {
			log.Fatal(err)
		}
		headerMap[name+"."+strconv.Itoa(i)] = data[i:10] //这个切片会导致整个底层数组被锁定,无法释放,如果底层数组较大
		//会对内存产生很大的压力
		headerMap[name+"."+strconv.Itoa(i)] = append([]byte{}, data[i:10]...) //将结果克隆一份,可以释放底层的数组
	}
	for k, v := range headerMap {
		fmt.Println("k,v:", k, v)
	}
}

/***
空指针和空接口不等价,比如返回了一个错误指针,但是并不是空的error接口

func returnsError() error{

	var p *MyError = nil
	if bad(){
		p = ErrBad
	}
	return p //will always return a non-mil error.
}
*/
/**
内存地址会变化.
Go语言中对象的地址可能发生变化,因此指针不能从其他非指针类型的值生成.
当内存发生变化的时候,相关的指针会同步更新,但是非指针类型的变量不会做同步更新.
同理CGO中也不能保存Go对象地址.
*/
/***
Goroutine泄露
Go语言是带内存自动回收的特性,因此内存一般不会泄漏.但是Goroutine确实存在泄漏的怦,同时泄漏的Goroutine引用的内存无法被回收.
*/
//这段代码没懂,以后再仔细看,先记录下来
func GoroutineLeak() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := func(ctx context.Context) <-chan int {
		channel := make(chan int)
		go func() {
			for i := 0; ; i++ {
				select {
				case <-ctx.Done():
					recover()
				case channel <- i:
				}

			}
		}()
		return channel
	}(ctx)
	for v := range ch {
		println(v)
		if v == 5 {
			cancel()//当main函数在break跳出循环时,通过调用cancel()来通知后台Goroutine退出,就避免了Goroutine泄漏.
			break
		}
	}
}

func main() {
	//VarParam() //可变参数是空接口类型.
	//ArrayTransform()
	//MapIterator()
	//err := errors.New("return value mask")
	//ReturnValueMask(err)
	//RecoverMustBePerformInDeferFunction()
	//go MainEarlyTermination()
	//通过休眠Sleep来回避并发中的问题,但休眠并不能保证输出完整的字符串
	//time.Sleep(time.Second)
	//或者插入调度语句
	//runtime.Gosched()

	//MonopolizeCPU()
	//go setup()
	//<-done
	//println(msg)

	//ClosureError()
	//ClosureRight()

	SliceLockArray()
}
