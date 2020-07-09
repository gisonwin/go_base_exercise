package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

/**
*  @Author <a href="mailto:gisonwin@qq.com">GisonWin</a>
*  @Date  2019/8/19 9:19
 */

func sayhelloName(w http.ResponseWriter,r *http.Request){
	r.ParseForm() //解析参数,默认是不会解析的
	fmt.Println(r.Form)//输出到服务端的打印信息
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form {
		fmt.Println("key= ",k ," val= ",strings.Join(v," "))
	}
	fmt.Fprintf(w,"Hello Gisonwin")
}
/***
web工作方式的几个概念:
Request:用户请求的信息,用来解析用户的请求信息,包括post,get,cookie,url等信息
Response:服务器要反馈给客户端的信息
Conn:用户的每次请求链接
Handler:处理请求和生成返回信息的处理逻辑.

http包执行流程:
1.创建Listen Socket,监听指定端口,等待客户端请求到来
2.Listen Socket接受客户端的请求,得到Client Socket,接下来通过Client Socket与客户端通信
3.处理客户端的请求,首先从Client Socket读取HTTP请求的协议,如果是POST方法,还要读取客户端提交的数据,
然后交给相应的Handler处理请求,handler处理完毕准备好客户端需要的数据,通过client socket写给客户端.

在这整个过程我们只要发解清楚下面的三个问题就知道Go如何让web运行起来了
1.如何监听端口
2.如何接收客户端请求
3.如何分配handler
ListenAndServe底层是这样处理的:初始化一个server对象,然后调用了net.Listen("tcp",addr),也就是底层用
TCP协议搭建了一个服务,然后监控我们设置的端口.
 */
func main() {
	http.HandleFunc("/",sayhelloName)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatal("listenandserve:",err)
	}


	/**
	 Go代码的执行流程:
	首先调用http.HandleFunc
	按顺序做了几件事:
	1.调用了DefaultServerMux的HandleFunc
	2.调用了DefaultServerMux的Handle
	3.往DefaultServeMux的map[string]muxEntry中增加对应的Handler和路由规则
	其次调用了http.ListenAndServe(":9090",nil)
	按顺序做了几件事
	1.实例化Server
	2.调用Server的ListenAndServe()
	3.调用net.Listen("tcp",addr)监听端口
	4.启动一个for循环,在循环体中Accept请求
	5.对每个请求实例化一个Conn,并且开启一个goroutine为这个请求进行服务 go c.server()
	6.读取每个请求的内容w,err := c.readRequest()
	7.判断handler是否为空,如果没有设置handler,handler就设置为DefaultServeMux
	8.调用handler的ServeHttp
	9.进入到DefaultServerMux.ServeHttp
	10.根据request选择handler,并且进入到这个Handler的ServeHttp
	mux.handler(r).ServeHTTP(w,r)
	11.选择handler:
	A 判断是否有路由能满足这个request(循环遍历ServerMux的muxEntry)
	B 如果有路由满足,调用这个路由Handler的ServeHttp
	C 如果没有路由满足,调用NotFoundHandler的ServeHttp
	 */
}